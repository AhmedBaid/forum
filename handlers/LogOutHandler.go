package handlers

import (
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}
	if exists, _ := helpers.SessionChecked(w, r); !exists {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err = utils.Db.Exec("Update users set session = ? where session = ?", "Null", cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	expiredCookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &expiredCookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

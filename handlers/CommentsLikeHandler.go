package handlers

import (
	"database/sql"
	"net/http"

	"forum/helpers"
	"forum/utils"
)

func CommentsLikeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}
	_, session := helpers.SessionChecked(w, r)

	commentId := r.FormValue("commentID")
	userId := r.FormValue("userId")
	reaction := r.FormValue("reaction")
	if userId == "" || commentId == "" || reaction == "" || (reaction != "1" && reaction != "-1") {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return
	}

	stmt2 := "SELECT id  FROM users WHERE session = ?"
	var userid int
	errr := utils.Db.QueryRow(stmt2, session).Scan(&userid)

	if errr != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	stmt := `select value from commentsLikes
	where commentsLikes.commentID = ? and commentsLikes.userID = ?`

	row := utils.Db.QueryRow(stmt, commentId, userid)
	var reactionValue string
	err := row.Scan(&reactionValue)
	if err != nil {
		if err == sql.ErrNoRows {
			stmt := `insert into commentsLikes (commentID, userID, value) values(?, ?, ?)`
			_, err := utils.Db.Exec(stmt, commentId, userid, reaction)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
				return
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
			return
		}
	} else {
		if reactionValue == reaction {
			// delete the like
			stmt := `delete from commentsLikes where commentID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, commentId, userid)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
				return

			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {

			// update the like
			stmt := `update commentsLikes set value = ? where commentID = ? and userID = ?`
			_, err := utils.Db.Exec(stmt, reaction, commentId, userid)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)

				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

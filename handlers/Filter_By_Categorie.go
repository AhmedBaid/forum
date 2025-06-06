package handlers

import (
	"net/http"
	"time"

	"forum/helpers"
	"forum/utils"
)

func Filter_By_Categorie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusMethodNotAllowed, utils.ErrorMethodnotAll)
		return
	}
	err := r.ParseForm()
	if err != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusBadRequest, utils.ErrorBadReq)
		return

	}
	commentMap, err := helpers.FetchComments(r)

	categories, errcat := helpers.AllCategories()

	mapp, errall := helpers.FetchCategories()
	if err != nil || errcat != nil || errall != nil {
		helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, utils.ErrorInternalServerErr)
		return
	}

	cookie, errr := r.Cookie("session")
	var sessValue string
	if errr != nil {
		sessValue = ""
	} else {
		sessValue = cookie.Value
	}

	categ := r.Form["tags"]
	if len(categ) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	query := `  SELECT  
				p.id, 
				p.username, 
				p.title, 
				p.description, 
				p.time,
				COALESCE((p.image_path), '') AS imaghe_path,
				COUNT(CASE WHEN l.value = 1 THEN 1 END) AS total_likes, 
				COUNT(CASE WHEN l.value = -1 THEN 1 END) AS total_dislikes
				FROM posts p
				INNER JOIN categories_post cp ON p.id = cp.postID
				INNER JOIN categories c ON cp.categoryID = c.id
				LEFT JOIN likes l ON p.id = l.postID
				WHERE c.id = ?
				GROUP BY p.id
				ORDER BY p.time DESC
`
	var posts []utils.Posts
	var post utils.Posts
	mapp1 := make(map[int]bool) // pour éviter duplication
	var totalLikes, totalDislikes int

	for _, categorie := range categ {
		rows, err := utils.Db.Query(query, categorie)
		if err != nil {
			helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
			return
		}
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.Username, &post.Title, &post.Description, &post.Time, &post.ImagePath, &totalLikes, &totalDislikes)
			if err != nil {
				helpers.RanderTemplate(w, "statusPage.html", http.StatusInternalServerError, nil)
				return
			}
			if !mapp1[post.Id] {
				post.Comments = commentMap[post.Id]
				post.TotalLikes = totalLikes
				post.TotalDislikes = totalDislikes
				post.TotalComments = len(commentMap[post.Id])
				now := time.Now()
				diff := now.Sub(post.Time)
				seconds := int(diff.Seconds())
				post.TimeFormatted = helpers.FormatDuration((seconds))
				post.Categories = mapp[post.Id]
				posts = append(posts, post)
				mapp1[post.Id] = true
			}
		}
	}

	variables := struct {
		Session    string
		UserActive string
		Posts      []utils.Posts
		Categories []utils.Categories
	}{
		Session:    sessValue,
		UserActive: helpers.GetUsernameFromSession(sessValue),
		Posts:      posts,
		Categories: categories,
	}

	helpers.RanderTemplate(w, "home.html", 200, variables)
}

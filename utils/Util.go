package utils

import (
	"database/sql"
	"net/http"
	"text/template"
)

var (
	Tp *template.Template
	Db *sql.DB
)

type ErrorPage struct {
	Code         int
	ErrorMessage string
}
type Users struct {
	Username string
	Email    string
	Password string
}
type Categories struct {
	name []string
}
type Posts struct {
	Username    string
	Title       string
	Description string
	Time        string
}




var (
	ErrorBadReq = ErrorPage{
		Code:         http.StatusBadRequest,
		ErrorMessage: "Oops! It looks like there was an issue with your request. Please check your input and try again.",
	}

	ErrorNotFound = ErrorPage{
		Code:         http.StatusNotFound,
		ErrorMessage: "Uh-oh! The page you're looking for doesn't exist. It might have been moved or deleted.",
	}

	ErrorMethodnotAll = ErrorPage{
		Code:         http.StatusMethodNotAllowed,
		ErrorMessage: "The request method is not supported for this resource. Please check and try again with a valid method.",
	}

	ErrorInternalServerErr = ErrorPage{
		Code:         http.StatusInternalServerError,
		ErrorMessage: "Something went wrong on our end. We're working on fixing it—please try again later!",
	}
)

package main

import (
	"database/sql"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router
var db *sql.DB

func main() {
	database.Initialize()
	db = database.DB
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	http.ListenAndServe(":2000", middlewares.RemoveTrailingSlash(router))
}

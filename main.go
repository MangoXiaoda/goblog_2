package main

import (
	"goblog_2/app/http/middlewares"
	"goblog_2/bootstrap"
	"goblog_2/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

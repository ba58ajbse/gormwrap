package main

import (
	"gormwrap/infrastructure/database"
	"gormwrap/infrastructure/router"

	"github.com/gin-gonic/gin"
)

type App struct {
	engin *gin.Engine
	sql   *database.SQLHandler
}

var app *App

func main() {

	router.NewRouter(app.engin, app.sql)

	app.engin.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

func init() {
	app = &App{
		engin: gin.Default(),
		sql:   database.Init(),
	}
}

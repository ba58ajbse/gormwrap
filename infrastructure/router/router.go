package router

import (
	"gormwrap/handlers"
	"gormwrap/infrastructure/database"
	"gormwrap/infrastructure/repository"
	"gormwrap/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, sql *database.SQLHandler) {
	createUserHandler := buildGetUsers(sql)

	public := r.Group("/")

	public.GET("/_health", handlers.HealthCheck)
	public.POST("/user", createUserHandler)
}

func buildGetUsers(sql *database.SQLHandler) func(c *gin.Context) {
	repo := repository.NewUserRepository(*sql)
	uc := usecase.NewUserUsecase(repo)
	h := handlers.NewUserHandler(uc)

	return h.CreateUser
}

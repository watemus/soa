package handler

import (
	"github.com/gin-gonic/gin"
	"main-service/internal/repository"
	"main-service/internal/service"
)

type Handler struct {
	auth  *service.AuthService
	tasks *repository.Tasks
}

func NewHandler(auth *service.AuthService, tasks *repository.Tasks) *Handler {
	return &Handler{
		auth:  auth,
		tasks: tasks,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentify)
	{
		api.POST("/update-user", h.updateUser)
		api.POST("/create-task", h.createTask)
		api.POST("/update-task", h.getTask)
		api.POST("/delete-task", h.deleteTask)
		api.GET("/get-task", h.getTask)
		api.GET("/get-list-tasks", h.getListTasks)
	}
	return router
}

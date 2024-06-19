package handler

import (
	"github.com/gin-gonic/gin"
	"main-service/internal/repository"
	"main-service/internal/service"
	"net/http"
)

type Handler struct {
	auth   *service.AuthService
	tasks  *repository.Tasks
	events *service.Events
}

func NewHandler(auth *service.AuthService, tasks *repository.Tasks, events *service.Events) *Handler {
	return &Handler{
		auth:   auth,
		tasks:  tasks,
		events: events,
	}
}

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{})
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/health-check", h.healthCheck)
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
		api.POST("/send-like", h.sendLike)
		api.POST("/send-view", h.sendView)
	}
	return router
}

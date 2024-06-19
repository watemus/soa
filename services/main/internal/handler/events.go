package handler

import (
	"github.com/gin-gonic/gin"
	"main-service/internal/repository"
	"net/http"
)

type EventRequest struct {
	TaskId uint64 `json:"task_id"`
}

type EventResponse struct {
	Status string `json:"status"`
}

func (h *Handler) sendLike(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &EventRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	event := repository.StatEvent{
		Username: username,
		TaskId:   body.TaskId,
	}
	err = h.events.SendLike(event)
	if err != nil {
		newErrorResponse(c, "error while producing event", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, &EventResponse{
		Status: "ok",
	})
}

func (h *Handler) sendView(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &EventRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	event := repository.StatEvent{
		Username: username,
		TaskId:   body.TaskId,
	}
	err = h.events.SendView(event)
	if err != nil {
		newErrorResponse(c, "error while producing event", http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, &EventResponse{
		Status: "ok",
	})
}

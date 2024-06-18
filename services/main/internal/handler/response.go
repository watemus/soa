package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(c *gin.Context, msg string, statusCode int) {
	c.AbortWithStatusJSON(statusCode, errorResponse{
		Error: msg,
	})
}

type statusResponse struct {
	Status string `json:"status"`
}

func newStatusResponse(c *gin.Context, msg string, statusCode int) {
	c.JSON(statusCode, &statusResponse{
		Status: msg,
	})
}

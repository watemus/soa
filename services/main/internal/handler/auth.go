package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main-service/internal/server"
	"main-service/internal/service"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "username"
)

type SignUpBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) getUsername(c *gin.Context) (string, error) {
	username, found := c.Get(userCtx)
	if !found {
		return "", errors.New("user not found in ctx")
	}
	return username.(string), nil
}

func (h *Handler) signUp(c *gin.Context) {
	body := &SignUpBody{}
	err := c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(body)
	err = h.auth.SignUp(body.Username, body.Password)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	newStatusResponse(c, "ok", http.StatusOK)
}

func (h *Handler) signIn(c *gin.Context) {
	body := &SignInBody{}
	err := c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.auth.SignIn(body.Username, body.Password)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, "empty auth header", http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, "invalid auth header", http.StatusUnauthorized)
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, "token is empty", http.StatusUnauthorized)
		return
	}

	username, err := service.ParseJWT(headerParts[1])
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	c.Set(userCtx, username)
}

func (h *Handler) updateUser(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}

	body := &server.UpdateUserBody{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.auth.UpdateUser(username, body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	newStatusResponse(c, "ok", http.StatusOK)
}

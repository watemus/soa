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

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body SignUpBody true "account info"
// @Success 200 {integer} string "status"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body SignInBody true "credentials"
// @Success 200 {object} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
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
	log.Printf("username: %s", username)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}

	c.Set(userCtx, username)
}

// @Summary UpdateUser
// @Tags auth
// @Description update user data
// @ID login
// @Accept  json
// @Produce  json
// @Param input body server.UpdateUserBody true "params"
// @Success 200 {object} string "status"
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/update-user [post]
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

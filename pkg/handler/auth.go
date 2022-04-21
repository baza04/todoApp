package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	todoapp "github.com/baza04/todoApp"
)

// @Summary Sign Up (Registration)
// @Description return created user ID
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body todoapp.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var (
		input todoapp.User
		id    int
		err   error
	)

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if id, err = h.services.Authorization.CreateUser(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign In (Authentication)
// @Description return token
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body signInInput true "username password"
// @Success 200 {string} string token "someToken"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var (
		input signInInput
		token string
		err   error
	)

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if token, err = h.services.Authorization.GenerateToken(input.Username, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

package handler

import (
	"Advertising/pkg/errormessage"
	"Advertising/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		errormessage.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(user)
	if err != nil {
		errormessage.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignIn struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignIn

	if err := c.BindJSON(&input); err != nil {
		errormessage.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.GenerateToken(input.Login, input.Password)
	if err != nil {
		errormessage.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

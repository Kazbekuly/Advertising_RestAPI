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

func (h *Handler) signIn(c *gin.Context) {

}

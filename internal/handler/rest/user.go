package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {

}

func (h *Handler) getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (h *Handler) getUserByID(c *gin.Context) {

}

func (h *Handler) updateUser(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}

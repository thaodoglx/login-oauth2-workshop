package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout controller for logout flow from system
func Logout(c *gin.Context) {
	c.Redirect(http.StatusFound, "http://127.0.0.1:9000/.ory/kratos/public/self-service/browser/flows/logout")
}

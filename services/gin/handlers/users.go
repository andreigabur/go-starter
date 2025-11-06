package handlers

import (
	"gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUsers(c *gin.Context) {

	users := controllers.ListUsers()

	c.IndentedJSON(http.StatusOK, users)
}

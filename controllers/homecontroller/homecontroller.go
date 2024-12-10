package homecontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	// Render the home page without passing any data
	c.HTML(http.StatusOK, "home/index.html", nil)
}

package Routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("api/userapi")
	{
		grp1.GET("user", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello chatroom")
		})
	}

	return r
}

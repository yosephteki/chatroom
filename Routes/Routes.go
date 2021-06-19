package Routes

import (
	"chatroom/Controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp1 := r.Group("api/chat")
	{
		grp1.GET("user", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello chatroom")
		})
		grp1.GET("Conversation/:id", Controllers.GetAllChatByUser)
	}

	return r
}

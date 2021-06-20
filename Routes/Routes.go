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
		grp1.GET("welcome", func(c *gin.Context) {
			c.JSON(http.StatusOK, "Hello chatroom")
		})
		grp1.GET("conversation/:id", Controllers.GetAllConversationByUser)
		grp1.GET("conversation/chatroom/:sender/:recipient", Controllers.GetConversation)
		grp1.POST("/:sender/:recipient/:message", Controllers.SendChatToUser)
	}

	return r
}

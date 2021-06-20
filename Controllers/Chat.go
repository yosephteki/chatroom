package Controllers

import (
	"chatroom/Models"
	"chatroom/Usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllConversationByUser(c *gin.Context) {
	var conversation []Models.ShowConversation
	userid := c.Params.ByName("id")

	err := Usecase.GetAllConversationByUser(&conversation, userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Conversation not found")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, conversation)
	}

}

func SendChatToUser(c *gin.Context) {
	sender := c.Params.ByName("sender")
	recipient := c.Params.ByName("recipient")
	message := c.Params.ByName("message")

	err := Usecase.SendChatToUser(sender, recipient, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Fail to send a chat",
				"error":   err,
			})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, "send message success!")
	}

}

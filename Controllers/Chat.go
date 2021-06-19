package Controllers

import (
	"chatroom/Models"
	"chatroom/Usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllChatByUser(c *gin.Context) {
	var chats []Models.Chat
	userid := c.Params.ByName("id")

	err := Usecase.GetAllConversationByUser(&chats, userid)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Conversation not found")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, chats)
	}

}

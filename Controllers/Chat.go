package Controllers

import (
	"chatroom/Models"
	"chatroom/Usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllConversationByUser(c *gin.Context) {
	var conversation []Models.ShowConversation
	userid := c.Params.ByName("id")

	err := Usecase.GetAllConversationByUser(&conversation, userid)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Conversation not found")
		c.Abort()
	} else {
		c.JSON(http.StatusOK, conversation)
	}

}

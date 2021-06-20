package Controllers

import (
	"chatroom/Models"
	"chatroom/Usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- task 4 ---
// User can list all their conversations (if user A has been chatting with user C & D, the list for A will shows A-C & A-D)
// explanation : this endpoint will return list of conversation of user by inputing userid as parameter
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

// --- task 1 & 3 ---
// 1. User can send a message to another user.
// 3. User can reply to a conversation they are involved with.
// explanation : user can send message and reply using this function, by reverse the sender and recipient value(user id)
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

// --- task 2 ---
// User can list all messages in a conversation between them and another user
// explanation : this endpoint returning chatroom object that contain sender and recipient and array of chat
func GetConversation(c *gin.Context) {
	sender := c.Params.ByName("sender")
	recipient := c.Params.ByName("recipient")
	var conversation []Models.Conversation

	err := Usecase.GetConversation(sender, recipient, &conversation)
	fmt.Println(len(conversation))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "Fail fetch conversation",
				"error":   err,
			})
		c.Abort()
	} else {
		var chatroom Models.Chatroom
		chatroom.Participant1 = sender
		chatroom.Participant2 = recipient
		chatroom.Conversation = conversation
		c.JSON(http.StatusOK, chatroom)
	}

}

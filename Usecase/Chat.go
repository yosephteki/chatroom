package Usecase

import (
	"chatroom/Models"
	"chatroom/Repositories"
)

func GetAllConversationByUser(conversation *[]Models.ShowConversation, userid string) (err error) {
	err = Repositories.GetAllConversationByUser(conversation, userid)
	if err != nil {
		return err
	}
	return nil
}

func SendChatToUser(sender string, recipient string, message string) (err error) {
	err = Repositories.SendChatToUser(sender, recipient, message)
	if err != nil {
		return err
	}
	return nil
}

func GetConversation(sender string, recipient string, conversation *[]Models.Conversation) (err error) {
	err = Repositories.GetConversation(sender, recipient, conversation)
	if err != nil {
		return err
	}
	return nil
}

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

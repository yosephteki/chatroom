package Usecase

import (
	"chatroom/Models"
	"chatroom/Repositories"
)

func GetAllConversationByUser(chat *[]Models.Chat, userid string) (err error) {
	err = Repositories.GetAllConversationByUser(chat, userid)
	if err != nil {
		return err
	}
	return nil
}

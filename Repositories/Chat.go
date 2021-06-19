package Repositories

import (
	"chatroom/Config"
	"chatroom/Models"
)

func GetAllConversationByUser(chats *[]Models.Chat, userid string) (err error) {
	if err = Config.DB.Distinct("sender").Where("recipient = ?", userid).Find(&chats).Error; err != nil {
		return err
	}
	return nil
}

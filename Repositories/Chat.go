package Repositories

import (
	"chatroom/Config"
	"chatroom/Models"
)

func GetAllConversationByUser(conversation *[]Models.ShowConversation, userid string) (err error) {
	err = Config.DB.Raw(`select "user"."name" as "name" , x.sender from "user" left join ( select distinct(sender) from chat c where c.recipient = ?) x on x.sender = "user".id where x.sender is not null`, userid).Scan(&conversation).Error
	if err != nil {
		return err
	}

	// if err = Config.DB.Distinct("sender").Where("recipient = ?", userid).Find(&chats).Error; err != nil {
	// 	return err
	// }
	return nil
}

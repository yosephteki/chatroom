package Repositories

import (
	"chatroom/Config"
	"chatroom/Models"
)

func GetAllConversationByUser(conversation *[]Models.ShowConversation, userid string) (err error) {
	err = Config.DB.Raw(
		`select "user"."name" as "name" , x.sender,x.message
		from "user" 
		left join ( 
		select t2.sender,t2.created_at,t1.message 
		from chat t1
		inner join(
		select distinct(c2.sender),max(c2.created_at) as created_at 
		from chat c2
		where c2.recipient = ?
		group by c2.sender
		) t2 on t1.created_at = t2.created_at and t1.sender = t2.sender
		) x on x.sender = "user".id 
		where x.sender is not null`, userid).Scan(&conversation).Error
	if err != nil {
		return err
	}
	return nil
}

func SendChatToUser(sender string, recipient string, message string) (err error) {
	err = Config.DB.Exec(
		`INSERT INTO "chat" ("sender","recipient","message") 
		VALUES(?,?,?)`, sender, recipient, message).Error
	if err != nil {
		return err
	}
	return nil
}

func GetConversation(sender string, recipient string, conversation *[]Models.Conversation) (err error) {
	err = Config.DB.Raw(
		`select distinct(c.message),u."name",u.id as senderid,c.created_at 
		from "user" u 
		join( 
			select *
			from chat c 
			where c.recipient = ? and c.sender = ?
			or 
			c.sender = ? and c.recipient = ?
		) c on c.sender = u.id
		order by c.created_at asc`, recipient, sender, recipient, sender).Scan(&conversation).Error
	if err != nil {
		return err
	}
	return nil
}

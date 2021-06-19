package Models

type ChatRoom struct {
	Id     string `json:"id"`
	ChatId string `json:"chat_id"`
	UserId string `json:"user_id"`
}

func (u *ChatRoom) TableName() string {
	return "user"
}

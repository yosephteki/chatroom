package Models

type Chat struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (u *Chat) TableName() string {
	return "chat"
}

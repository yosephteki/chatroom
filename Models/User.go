package Models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Password string    `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

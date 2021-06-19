package Models

import (
	"github.com/google/uuid"
)

type Chat struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Sender    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Recipient uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Message   string    `json:"message"`
}

func (u *Chat) TableName() string {
	return "chat"
}

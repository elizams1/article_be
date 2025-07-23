package posts

import (
	"time"
)

type Posts struct {
	ID          int       `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" db:"title" validate:"required,min=20" gorm:"type:varchar(200);not null"`
	Content     string    `json:"content" db:"content" validate:"required,min=200" gorm:"type:text;not null"`
	Category    string    `json:"category" db:"category" validate:"required,min=3" gorm:"type:varchar(100);not null"`
	Status      string    `json:"status" db:"status" gorm:"type:enum('publish', 'draft', 'trash');default:'draft';not null"`
	CreatedDate time.Time `json:"created_date" db:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `json:"updated_date" db:"updated_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

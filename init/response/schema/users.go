package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `gorm:"column:id; primaryKey" json:"id"`
	Username  string `gorm:"column:username; type:varchar(255)" json:"username"`
	Password  string `gorm:"column:password; type:text" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package schema

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	Id        int  `gorm:"column:id; primaryKey" json:"id"`
	IdUser    int  `gorm:"column:id_user" json:"id_user"`
	User      User `gorm:"foreignKey:IdUser; references:id"`
	Amount    int  `gorm:"column:amount" json:"amount"`
	Tenor     int  `gorm:"column:tenor" json:"tenor"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

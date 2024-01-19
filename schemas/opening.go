package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct{
	gorm.Model
	Role string
	company string
	location string
	remote bool
	link string
	Salary int64
}

type OpeningResponse struct{
	ID uint `json:"id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt, omitempty"`
	Role string `json:"Role"`
	Company string `json:"Company"`
	Location string `json:"Location"`
	Remote bool `json:"Remote"`
	Link string `json:"Link"`
	Salary int64 `json:"Salary"`
}
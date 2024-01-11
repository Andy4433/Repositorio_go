package schemas

import (
	
	"gorm.io/gorm"
)

  type Opeing struct{
	gorm.Model
	Role string
	company string
	location string
	remote bool
	link string
	Salary int64
  }
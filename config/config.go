package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error{
	// return errors.New("test erro")
	return nil
}

func GetLogger(p string) *Logger{
	logger=NewLogger(p)
	return logger
}
package config

import (
	"os"
	"github.com/Andy4433/Repositorio_go.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//database creation process
func InitializeSQLite() (*gorm.DB, error) {
    logger := GetLogger("sqlite")

    dbPath := "./db/main.db"

    _, err := os.Stat(dbPath)
    if os.IsNotExist(err) {
        logger.Info("Database file not found, creating...")
        err = os.MkdirAll("./db", os.ModePerm)
        if err != nil {
            return nil, err
        }
        file, err := os.Create(dbPath)
        if err != nil {
            return nil, err
        }
        file.Close()
    }

    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        logger.Errorf("SQLite opening error: %v", err)
        return nil, err
    }

    err = db.AutoMigrate(&schemas.Opeing{})
    if err != nil {
        logger.Errorf("SQLite automigrate error: %v", err)
        return nil, err
    }

    return db, nil
}

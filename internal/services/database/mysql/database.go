package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"scaffold-game-nft-marketplace/pkg/config"
)

// returns a new database connector for the application.
func Connect(config config.DBConfig) (*gorm.DB, error) {
	// Set some mandatory parameters
	config.Params["parseTime"] = "true"
	config.Params["rejectReadOnly"] = "true"

	db, err := gorm.Open(mysql.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("connect database failed, error: %v", err))
	}

	return db, nil
}

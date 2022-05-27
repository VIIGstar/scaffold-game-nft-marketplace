package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"scaffold-game-nft-marketplace/pkg/config"
)

// Connect returns a new database connector for the application.
func Connect(config config.DBConfig) (*gorm.DB, error) {
	// Set some mandatory parameters
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.Name)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	pDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// check db
	err = pDB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

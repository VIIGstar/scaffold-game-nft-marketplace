package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"logur.dev/logur"
	"scaffold-game-nft-marketplace/internal/entities"
	_ "scaffold-game-nft-marketplace/internal/entities"
	"scaffold-game-nft-marketplace/internal/services"
	"scaffold-game-nft-marketplace/internal/services/database"
	"scaffold-game-nft-marketplace/internal/services/log"
	"scaffold-game-nft-marketplace/pkg/config"
)

type MigrateService struct {
	services.DefaultService
	logger logur.LoggerFacade
	db     *database.DB
}

func main() {
	migrateService := MigrateService{}
	migrateService.Init()
	defer migrateService.db.Close()

	tables := []interface{}{
		entities.Investor{},
		entities.User{},
		entities.Category{},
		entities.Asset{},
		entities.AssetTransaction{},
	}

	err := migrateService.db.GormDB().AutoMigrate(tables...)
	if err != nil {
		migrateService.logger.Error(fmt.Sprintf("Seed failed, details: %v", err))
		return
	}

	migrateService.logger.Info("Seed completed")
}

func (s *MigrateService) Init() {
	s.DefaultService.Init()
	var (
		logCf = config.LogConfig{}
		dbCf  = config.DBConfig{}
	)
	cfBytes, _ := json.Marshal(viper.GetStringMap("log"))
	json.Unmarshal(cfBytes, &logCf)
	cfBytes, _ = json.Marshal(viper.GetStringMap("database"))
	json.Unmarshal(cfBytes, &dbCf)

	logger := log.NewLogger(logCf)

	// Override the global standard library logger to make sure everything uses our logger
	log.SetStandardLogger(logger)
	// Start database
	if dbCf.Params == nil {
		dbCf.Params = make(map[string]string)
	}

	s.db = database.New(logger, dbCf)
	s.logger = logger
}

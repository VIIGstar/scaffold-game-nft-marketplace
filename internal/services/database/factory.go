package database

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	mysql "scaffold-game-nft-marketplace/internal/services/database/mysql"
	"scaffold-game-nft-marketplace/internal/services/database/postgres"
	"scaffold-game-nft-marketplace/pkg/config"
)

func NewConnector(cfg config.DBConfig) (*gorm.DB, error) {
	if viper.GetBool("use_mysql") {
		return mysql.Connect(cfg)
	}

	return postgres.Connect(config.DBConfig{
		Host:   viper.GetString("psql.host"),
		Port:   viper.GetInt("psql.port"),
		User:   viper.GetString("psql.user"),
		Pass:   viper.GetString("psql.pass"),
		Name:   viper.GetString("psql.name"),
		Params: map[string]string{},
	})
}

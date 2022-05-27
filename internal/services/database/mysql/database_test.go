package database

import (
	health "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/stretchr/testify/assert"
	"scaffold-game-nft-marketplace/internal/services/health-check"
	"scaffold-game-nft-marketplace/internal/services/log"
	"scaffold-game-nft-marketplace/pkg/config"
	"testing"
	"time"
)

func TestConfig_Connector(t *testing.T) {
	cf := config.DBConfig{
		Host:   "localhost",
		Port:   3306,
		User:   "root",
		Pass:   "",
		Name:   "dev",
		Params: make(map[string]string),
	}
	lg := log.NewLogger(config.LogConfig{})
	SetLogger(lg)
	gormDB, err := NewConnector(cf)
	assert.Equal(t, nil, err)
	db, _ := gormDB.DB()
	defer db.Close()

	healthListener := health_check.NewHealthListener(lg, "mysql")
	healthChecker := health.New(health.WithHealthListeners(healthListener))
	// Register database health check
	_ = healthChecker.RegisterCheck(
		checks.Must(checks.NewPingCheck("db.check", db)),
		health.ExecutionPeriod(time.Millisecond*10))

	time.Sleep(time.Millisecond * 20)
}

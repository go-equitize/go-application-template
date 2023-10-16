package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConfig(t *testing.T) {
	cfg := DBConfig{
		Host:        "localhost",
		Port:        3306,
		DBName:      "tx_aggregator",
		User:        "user",
		Password:    "password",
		ConnTimeOut: 30,
	}
	assert.Equal(t, "user:password@tcp(localhost:3306)/tx_aggregator?charset=utf8mb4&parseTime=True&timeout=30s", cfg.DSN())
	assert.Equal(t, "user:password@tcp(localhost:3306)/tx_aggregator?multiStatements=true", cfg.MigrationSource())
}

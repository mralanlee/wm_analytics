package main

import (
	"fmt"
	"os"

	"github.com/mralanlee/wm_analytics/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	PostgresUser string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort string
	DSN string
}

var PostgresClient *gorm.DB

var (
	postgres_user = os.Getenv("POSTGRES_USER")
	postgres_pass = os.Getenv("POSTGRES_PASSWORD")
	postgres_db = os.Getenv("POSTGRES_DBNAME")
	postgres_port = os.Getenv("POSTGRES_PORT")
)

func (c *PostgresConfig) fill() {
	if c.PostgresUser == "" {
		c.PostgresUser = common.POSTGRES_USER
	}

	if c.PostgresPassword == "" {
		c.PostgresPassword = common.POSTGRES_PASS
	}

	if c.PostgresDatabase == "" {
		c.PostgresDatabase = common.POSTGRES_DBNAME
	}

	if c.PostgresPort == "" {
		c.PostgresPort = common.POSTGRES_PORT
	}

	c.DSN = fmt.Sprintf("user=%s password=%s dbname=%s port=%s", c.PostgresUser, c.PostgresPassword, c.PostgresDatabase, c.PostgresPort)
}

func init() {
	var config = &PostgresConfig{
		PostgresUser: postgres_user,
		PostgresPassword: postgres_pass,
		PostgresDatabase: postgres_db,
		PostgresPort: postgres_port,
	}

	config.fill()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	PostgresClient = db
}

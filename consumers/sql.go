package consumers

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
	PostgresHost string
	PostgresPort string
	DSN string
}

var PostgresClient *gorm.DB

var (
	postgres_user = os.Getenv("POSTGRES_USER")
	postgres_pass = os.Getenv("POSTGRES_PASSWORD")
	postgres_db = os.Getenv("POSTGRES_DBNAME")
	postgres_host = os.Getenv("POSTGRES_HOST")
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

	if c.PostgresHost == "" {
		c.PostgresHost = common.POSTGRES_HOST
	}

	if c.PostgresPort == "" {
		c.PostgresPort = common.POSTGRES_PORT
	}

	c.DSN = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", c.PostgresUser, c.PostgresPassword, c.PostgresDatabase, c.PostgresHost, c.PostgresPort)
}

func init() {
	var config = &PostgresConfig{
		PostgresUser: postgres_user,
		PostgresPassword: postgres_pass,
		PostgresDatabase: postgres_db,
		PostgresHost: postgres_host,
		PostgresPort: postgres_port,
	}

	config.fill()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&PaymentDetail{})

	PostgresClient = db
}

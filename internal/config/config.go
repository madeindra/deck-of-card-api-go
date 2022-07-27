package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App struct {
		Port string
	}
	Database struct {
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
	}
}

func Init() *Config {
	c := new(Config)
	c.initApp()
	c.initDB()

	return c
}

// configuration for app
func (c *Config) initApp() *Config {
	port := os.Getenv("APP_PORT")
	c.App.Port = port

	return c
}

// configuration for database instance
func (c *Config) initDB() *Config {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	maxOpenConnections, _ := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNECTIONS"), 10, 64)
	maxIdleConnections, _ := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONNECTIONS"), 10, 64)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, database)

	c.Database.DSN = dsn

	c.Database.MaxOpenConnections = int(maxOpenConnections)
	c.Database.MaxIdleConnections = int(maxIdleConnections)

	return c
}

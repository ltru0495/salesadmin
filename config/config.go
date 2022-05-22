package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var database *DatabaseConfig
var server *ServerConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USERNAME", "root")
	database.Password = gonv.GetStringEnv("PASSWORD", "")
	database.Host = gonv.GetStringEnv("HOST", "127.0.0.1")
	database.Port = gonv.GetIntEnv("PORT", 27027)
	database.Database = gonv.GetStringEnv("DATABASE", "salesadmin")

	server = &ServerConfig{}
	server.host = gonv.GetStringEnv("HOST", "127.0.0.1")
	server.port = gonv.GetIntEnv("PORT", 9000)
	server.debug = gonv.GetBoolEnv("DEBUG", true)
}

func DirTemplate() string {
	return "templates/*.html"
}

func DirTemplateError() string {
	return "templates/error.html"
}

func GetDatabaseConfig() *DatabaseConfig {
	return database
}

func UrlServer() string {
	return server.url()
}

func ServerPort() string {
	return fmt.Sprintf(":%d", server.port)
}

func (this *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", this.host, this.port)
}

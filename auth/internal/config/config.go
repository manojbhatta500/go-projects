package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SslMode  string
	}
	Jwt struct {
		SecretKey string
	}
	Enviroment string
}

var ConfigInstance Config

func (c *Config) GetDbConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.SslMode,
	)
}

func (c *Config) ConfigPrinter() {
	res := fmt.Sprintf("server  : host  : %s : port : %s", c.Server.Host, c.Server.Port)
	fmt.Println(res)
	dbres := fmt.Sprintf("database: host  %s  , port  %s , user %s  , password  %s ,dbname  %s ,sslmode %s", c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBName, c.Database.SslMode)
	fmt.Println(dbres)
	fmt.Println(c.Enviroment)

}

func Load() error {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("something went wrong while loading .env file", err.Error())

		return err
	}

	ConfigInstance.Server.Port = os.Getenv("SERVER_PORT")
	ConfigInstance.Server.Host = os.Getenv("SERVER_HOST")
	ConfigInstance.Database.Host = os.Getenv("DB_HOST")
	ConfigInstance.Database.Port = os.Getenv("DB_PORT")
	ConfigInstance.Database.User = os.Getenv("DB_USER")
	ConfigInstance.Database.Password = os.Getenv("DB_PASSWORD")
	ConfigInstance.Database.SslMode = os.Getenv("DB_SSLMODE")
	ConfigInstance.Jwt.SecretKey = os.Getenv("JWT_SECRET")
	ConfigInstance.Enviroment = os.Getenv("ENV")

	return nil

}

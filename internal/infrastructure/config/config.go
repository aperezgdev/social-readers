package config

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Environment string

const (
	DEVELOP    Environment = "DEV"
	PRODUCTION Environment = "PROD"
)

type Config struct {
	ServerPort       string
	Env              Environment
	DatabaseUrl      string
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
}

func NewConfig(slog *slog.Logger) Config {
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	readTimeoutN, errRead := strconv.ParseInt(os.Getenv("SERVER_READ_TIMEOUT"), 10, 64)
	if errRead != nil {
		slog.Error("NewConfig - Server Read Timeout is not a valid number")
		panic(errRead)
	}

	writeTimeoutN, errWrite := strconv.ParseInt(os.Getenv("SERVER_WRITE_TIMEOUT"), 10, 64)
	if errWrite != nil {
		slog.Error("NewConfig - Server Write Timeout is not a valid number")
	}

	return Config{
		ServerPort:       os.Getenv("SERVER_PORT"),
		Env:              Environment(os.Getenv("ENV")),
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabaseHost:     os.Getenv("DB_HOST"),
		DatabaseUrl:      "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME"),
		DatabaseUser:     os.Getenv("DB_USER"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		ReadTimeout:      time.Duration(readTimeoutN) * time.Second,
		WriteTimeout:     time.Duration(writeTimeoutN) * time.Second,
	}
}

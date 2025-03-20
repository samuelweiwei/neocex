package config

import (
	"strings"

	"gorm.io/gorm/logger"
)

type DSN interface {
	Dsn() string
}

type GeneralDB struct {
	Prefix      string
	Port        string
	Config      string
	Dbname      string
	Username    string
	Password    string
	Sslmode     string
	Path        string
	Engine      string
	LogMode     string
	MaxIdleConn int
	MaxOpenConn int
	Singular    bool
	LogZap      bool
}

func (c GeneralDB) LogLevel() logger.LogLevel {
	switch strings.ToLower(c.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}

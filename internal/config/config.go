package config

import (
	"fmt"
	"time"

	"log/slog"

	"github.com/jessevdk/go-flags"
)

// LogLevel — кастомный тип для поддержки парсинга строк в slog.Level.
type LogLevel slog.Level

// UnmarshalFlag реализует интерфейс flags.Unmarshaler для парсинга строк.
func (l *LogLevel) UnmarshalFlag(value string) error {
	switch value {
	case "debug":
		*l = LogLevel(slog.LevelDebug)
	case "info":
		*l = LogLevel(slog.LevelInfo)
	case "warn":
		*l = LogLevel(slog.LevelWarn)
	case "error":
		*l = LogLevel(slog.LevelError)
	default:
		return fmt.Errorf("неизвестный уровень логирования: %s", value)
	}
	return nil
}

type Config struct {
	Server struct {
		Port         int           `long:"port" env:"PORT" default:"8080" description:"Порт сервера"`
		ReadTimeout  time.Duration `long:"read-timeout" env:"READ_TIMEOUT" default:"10s" description:"Таймаут чтения"`
		WriteTimeout time.Duration `long:"write-timeout" env:"WRITE_TIMEOUT" default:"10s" description:"Таймаут записи"`
		IdleTimeout  time.Duration `long:"idle-timeout" env:"IDLE_TIMEOUT" default:"60s" description:"Таймаут простоя"`
	} `group:"server" namespace:"server" env-namespace:"SERVER"`
	LogLevel LogLevel `long:"log-level" env:"LOG_LEVEL" default:"info" description:"Уровень логирования (debug, info, warn, error)"`
}

func Load() (*Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				return nil, nil
			}
		}
		return nil, err
	}

	return &cfg, nil
}

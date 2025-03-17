package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jessevdk/go-flags"
)

// Кастомный тип для уровня логирования
type LogLevel slog.Level

// Реализуем интерфейс Unmarshaler для поддержки парсинга строк
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

func main() {
	cfg := Config{}

	parser := flags.NewParser(&cfg, flags.Default)

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				os.Exit(0)
			}
		}
		fmt.Printf("Ошибка при парсинге флагов: %v\n", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(cfg.LogLevel)}))

	// Создаем HTTP-сервер
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		Handler:      http.HandlerFunc(handleRequest), // Простой обработчик запросов
	}

	// Каналы для обработки сигналов и ошибок
	stopChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)

	// Запуск сервера в отдельной горутине
	go func() {
		logger.Info("Сервер запущен", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("ошибка при запуске сервера: %w", err)
		}
	}()

	// Ожидание сигналов завершения
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stopChan:
		logger.Info("Получен сигнал завершения, остановка сервера...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("Ошибка при остановке сервера", "error", err)
		}
		logger.Info("Сервер остановлен")
	case err := <-errChan:
		logger.Error("Ошибка при работе сервера", "error", err)
		os.Exit(1)
	}
}

// Простой обработчик запросов
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}

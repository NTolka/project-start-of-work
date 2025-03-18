package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/NTolka/project-start-of-work/internal/api/http"
	"github.com/NTolka/project-start-of-work/internal/config"
	"github.com/NTolka/project-start-of-work/internal/logger"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Ошибка загрузки конфигурации", "error", err)
		os.Exit(1)
	}

	// Инициализация логгера
	log := logger.New(slog.Level(cfg.LogLevel)) // Преобразуем LogLevel в slog.Level

	// Инициализация и запуск HTTP-сервера
	server := http.NewServer(cfg, log)
	go func() {
		if err := server.Start(); err != nil {
			log.Error("Ошибка при запуске сервера", "error", err)
			os.Exit(1)
		}
	}()

	// Ожидание сигналов завершения
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	// Остановка сервера
	if err := server.Stop(); err != nil {
		log.Error("Ошибка при остановке сервера", "error", err)
		os.Exit(1)
	}

	log.Info("Сервер остановлен")
}

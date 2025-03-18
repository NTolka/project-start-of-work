package logger

import (
	"log/slog"
	"os"
)

// Logger — это обёртка над slog.Logger для удобства использования.
type Logger struct {
	*slog.Logger
}

// New создаёт новый логгер с указанным уровнем логирования.
func New(level slog.Level) *Logger {
	return &Logger{
		Logger: slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level})),
	}
}

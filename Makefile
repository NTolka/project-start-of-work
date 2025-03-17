# Переменные
APP_NAME=myapp
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOMOD=$(GO) mod
BIN_DIR=bin
LOG_LEVEL=info

# Цели
all: build

# Сборка приложения
build:
	@echo "Сборка приложения..."
	$(GOBUILD) -o $(BIN_DIR)/$(APP_NAME) ./main.go

# Очистка
clean:
	@echo "Очистка..."
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

# Запуск приложения
run:
	@echo "Запуск приложения..."
	$(GO) run ./cmd/main.go --log-level=$(LOG_LEVEL)

# Запуск тестов
test:
	@echo "Запуск тестов..."
	$(GOTEST) -v ./...

# Запуск линтера
lint:
	@echo "Запуск линтера..."
	golangci-lint run

# Запуск с отладкой
debug:
	@echo "Запуск с отладкой..."
	$(GO) run ./cmd/main.go --log-level=debug

# Установка зависимостей
deps:
	@echo "Установка зависимостей..."
	$(GOMOD) tidy

# Помощь
help:
	@echo "Использование: make [цель]"
	@echo ""
	@echo "Цели:"
	@echo "  all       Сборка приложения (по умолчанию)"
	@echo "  build     Сборка приложения"
	@echo "  clean     Очистка проекта"
	@echo "  run       Запуск приложения"
	@echo "  test      Запуск тестов"
	@echo "  lint      Запуск линтера"
	@echo "  debug     Запуск с отладкой (уровень логирования debug)"
	@echo "  deps      Установка зависимостей"
	@echo "  help      Вывод справки"
# Компилируем Go‑пакет в исполняемый файл
build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

# Запускаем линтер для проверки кода
lint:
	golangci-lint run

# Автоматически исправляем ошибки, которые может исправить линтер
lint-fix:
	golangci-lint run --fix
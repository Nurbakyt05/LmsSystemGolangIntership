# Используем минимальный и ARM-совместимый образ Go
FROM golang:1.24-alpine

# Устанавливаем необходимые зависимости
RUN apk update && apk add --no-cache git gcc musl-dev

# Создаем рабочую директорию внутри контейнера
WORKDIR /app

# Кэшируем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем весь проект внутрь контейнера
COPY . .

# Собираем Go-приложение
RUN go build -o main .

# Пробрасываем порт 8080
EXPOSE 8080

# Команда по умолчанию при запуске контейнера
CMD ["./main"]

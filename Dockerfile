# 1) Билдим приложение
FROM golang:1.24-alpine AS builder

WORKDIR /app

# зависимости
COPY go.mod go.sum ./
RUN go mod download

# копируем весь код
COPY . .

# ставим инструменты для сборки С
RUN apk update && apk add --no-cache git gcc musl-dev

# собираем бинарь с уникальным именем
RUN CGO_ENABLED=0 GOOS=linux go build -o lms-system ./main/main.go

# 2) Финальный образ
FROM alpine:latest

# (по желанию) корневой рабочий каталог
WORKDIR /root/

# копируем только бинарь из builder
COPY --from=builder /app/lms-system .

# делаем его исполняемым
RUN chmod +x lms-system

# запускаем приложение
ENTRYPOINT ["./lms-system"]

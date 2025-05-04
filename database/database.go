package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect инициализирует и возвращает соединение с БД (синглтон)
func Connect() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		// Получаем параметры подключения из переменных окружения
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5433")
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "NURIK2005")
		dbname := getEnv("DB_NAME", "LmsDataBase")
		sslmode := getEnv("DB_SSLMODE", "disable")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			host, user, password, dbname, port, sslmode)

		// Настройка логгера GORM
		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Измените на logger.Silent для production
		}

		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			log.Fatalf("🚨 Failed to connect to database: %v", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("🚨 Failed to get database instance: %v", err)
			return
		}

		// Настройка пула соединений
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)

		fmt.Println("✅ Connected to database")
	})

	return db, err
}

// GetDB возвращает существующее соединение с БД
func GetDB() *gorm.DB {
	return db
}

// Close закрывает соединение с БД
func Close() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB: %v", err)
	}

	return sqlDB.Close()
}

// Вспомогательная функция для получения переменных окружения
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

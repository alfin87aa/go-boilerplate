package configs

import (
	"boilerplate/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host     string
	Port     uint64
	DBName   string
	User     string
	Password string
}

var _db *gorm.DB

func BuildDBConfig() *DBConfig {
	port, err := strconv.ParseUint(os.Getenv("DB_PORT"), 0, 64)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Enable color
		},
	)
	dsn := DbURL(BuildDBConfig())

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(
		&models.User{},
	)

	db.Session(&gorm.Session{PrepareStmt: true})

	dbSQL, _ := db.DB()

	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	_db = db
}

func GetDB() *gorm.DB {
	return _db
}

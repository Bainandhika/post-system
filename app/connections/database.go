package connections

import (
	"fmt"
	"log"
	"os"
	"post-system/app/configs"
	"post-system/lib/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, error) {
	dbConfig := configs.DB
	dsn := fmt.Sprintf("host=%s dbname=%s port=%d user=%s password=%s sslmode=disable",
		dbConfig.Host, dbConfig.Name, dbConfig.Port, dbConfig.User, dbConfig.Password)

	fileName := configs.App.LogPath + "/post-system.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gormLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      gormLogger,
	})
	if err != nil {
		return nil, err
	}

	err = dbConnection.AutoMigrate(&models.Post{}, &models.Tag{})
	if err != nil {
		return nil, err
	}

	return dbConnection, nil
}

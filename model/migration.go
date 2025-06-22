package model

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

func Migration(dbName string, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file error: %v", err)
	}

	switch dbName {
	case os.Getenv("DATABASE_NAME"):
		db.AutoMigrate(&Series{})
		db.AutoMigrate(&Article{})
		db.AutoMigrate(&Comment{})
	case os.Getenv("DATABASE_NAME2"):
		db.AutoMigrate(&DownloadArticle{})
	default:
		logrus.Fatal("error in migration function")
	}
}

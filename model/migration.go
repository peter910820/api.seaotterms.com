package model

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	galmodel "api.seaotterms.com/model/gal"
	teachmodel "api.seaotterms.com/model/teach"
)

func Migration(dbName string, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf(".env file error: %v", err)
	}

	switch dbName {
	case os.Getenv("DATABASE_NAME"):
		db.AutoMigrate(&teachmodel.Series{})
		db.AutoMigrate(&teachmodel.Article{})
		db.AutoMigrate(&teachmodel.Comment{})
	case os.Getenv("DATABASE_NAME2"):
		// db.AutoMigrate(&galmodel.DownloadArticle{})
		db.AutoMigrate(&galmodel.Article{})
	default:
		logrus.Fatal("error in migration function")
	}
}

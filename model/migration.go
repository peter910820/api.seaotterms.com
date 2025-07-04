package model

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"

	// galmodel "api.seaotterms.com/model/gal"
	blogmodel "api.seaotterms.com/model/blog"
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
		// db.AutoMigrate(&galmodel.User{})
		// db.AutoMigrate(&galmodel.Tag{})
		// db.AutoMigrate(&galmodel.Article{})
		// db.AutoMigrate(&galmodel.Log{})
		// db.AutoMigrate(&galmodel.TmpData{})
	case os.Getenv("DATABASE_NAME3"):
		db.AutoMigrate(&blogmodel.Tag{})
		db.AutoMigrate(&blogmodel.Article{})
	default:
		logrus.Fatal("error in migration function")
	}
}

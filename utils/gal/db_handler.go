package gal

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	model "api.seaotterms.com/model/gal"
)

func WriteTmpData(dataType string, dataContent string, expirationAt time.Time, db *gorm.DB) error {
	data := model.TmpData{
		Type:         dataType,
		Content:      dataContent,
		ExpirationAt: expirationAt,
	}

	err := db.Create(&data).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

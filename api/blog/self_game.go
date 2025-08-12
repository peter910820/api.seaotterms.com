package blog

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	model "api.seaotterms.com/model/galgame"
	utils "api.seaotterms.com/utils/blog"
)

type GameRecordForClient struct {
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	ReleaseDate time.Time `json:"releaseDate"`
	AllAges     bool      `json:"allAges"`
	EndDate     time.Time `json:"endDate"`
	Username    string    `json:"username"`
}

type GameRecordForUpdate struct {
	ReleaseDate time.Time
	EndDate     time.Time
	UpdateName  string
	UpdateTime  time.Time
}

// use game name to query single galgame data
func QueryGalgame(c *fiber.Ctx, db *gorm.DB) error {
	var data []model.GameRecord
	// URL decoding
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = db.Where("name = ?", name).First(&data).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "找不到Game資料", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	logrus.Info("Galgame單筆資料查詢成功")
	response := utils.ResponseFactory(c, fiber.StatusOK, "Galgame單筆資料查詢成功", &data)
	return c.Status(fiber.StatusOK).JSON(response)
}

// get all the galgame data for specify brand
func QueryGalgameByBrand(c *fiber.Ctx, db *gorm.DB) error {
	var data []model.GameRecord
	// URL decoding
	brand, err := url.QueryUnescape(c.Params("brand"))
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = db.Where("brand = ?", brand).Order("end_date DESC").Find(&data).Error
	if err != nil {
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	// if data not exist, retrun a empty struct
	logrus.Info("Galgame多筆資料查詢成功")
	response := utils.ResponseFactory[any](c, fiber.StatusOK, "Galgame多筆資料查詢成功", nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

// update single galgame data (develop)
func UpdateGalgameDevelop(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData GameRecordForClient
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}
	// URL decoding
	name, err := url.QueryUnescape(c.Params("name"))
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// gorm:"autoUpdateTime" can not update, so manual update update_time
	err = db.Model(&model.GameRecord{}).Where("name = ?", name).
		Select("release_date", "end_date", "update_name", "update_time").
		Updates(GameRecordForUpdate{
			ReleaseDate: clientData.ReleaseDate,
			EndDate:     clientData.EndDate,
			UpdateName:  clientData.Username,
			UpdateTime:  time.Now(),
		}).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "找不到Game資料", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}
	logrus.Infof("資料 %s 更新成功", name)
	response := utils.ResponseFactory[any](c, fiber.StatusOK, fmt.Sprintf("資料 %s 更新成功", name), nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

// insert data to galgame
func CreateGalgame(c *fiber.Ctx, db *gorm.DB) error {
	// load client data
	var clientData GameRecordForClient
	if err := c.BodyParser(&clientData); err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	data := model.GameRecord{
		Name:        clientData.Name,
		Brand:       clientData.Brand,
		ReleaseDate: clientData.ReleaseDate,
		AllAges:     clientData.AllAges,
		EndDate:     clientData.EndDate,
		InputName:   clientData.Username,
		UpdateName:  clientData.Username,
	}
	err := db.Create(&data).Error
	if err != nil {
		logrus.Error(err)
		response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	logrus.Infof("資料 %s 創建成功", clientData.Name)

	/* --------------------------------- */
	/* --------------------------------- */

	// update brand info
	var brandData model.BrandRecord

	err = db.Where("brand = ?", clientData.Brand).First(&brandData).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "找不到Game資料", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	annotation := "待攻略"
	if brandData.Completed+1 == brandData.Total {
		annotation = "制霸"
	}

	// gorm:"autoUpdateTime" can not update, so manual update update_time
	err = db.Model(&model.BrandRecord{}).Where("brand = ?", clientData.Brand).
		Select("completed", "annotation", "update_name", "update_time").
		Updates(BrandRecordForUpdate{
			Completed:  brandData.Completed + 1,
			Annotation: annotation,
			UpdateName: clientData.Username,
			UpdateTime: time.Now(),
		}).Error
	if err != nil {
		logrus.Error(err)
		// if record not exist
		if err == gorm.ErrRecordNotFound {
			response := utils.ResponseFactory[any](c, fiber.StatusNotFound, "找不到Game資料", nil)
			return c.Status(fiber.StatusNotFound).JSON(response)
		} else {
			response := utils.ResponseFactory[any](c, fiber.StatusInternalServerError, err.Error(), nil)
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
	}

	logrus.Infof("資料 %s 創建成功", clientData.Name)
	response := utils.ResponseFactory[any](c, fiber.StatusOK, fmt.Sprintf("資料 %s 創建成功", clientData.Name), nil)
	return c.Status(fiber.StatusOK).JSON(response)
}

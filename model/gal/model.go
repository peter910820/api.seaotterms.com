package gal

import (
	"time"
)

// A00_Galgame

// type DownloadArticle struct {
// 	ID           uint      `gorm:"primaryKey" json:"id"`
// 	Image        string    `json:"image"`
// 	Content      string    `gorm:"NOT NULL" json:"content"`
// 	DownloadType string    `gorm:"NOT NULL" json:"downloadType"`
// 	DownloadURL  string    `gorm:"NOT NULL" json:"downloadUrl"`
// 	CreatedAt    time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
// 	CreatedName  string    `gorm:"NOT NULL" json:"createdName"`
// 	UpdatedAt    time.Time `json:"updatedAt"`
// 	UpdatedName  string    `json:"updatedName"`
// }

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Email          string    `gorm:"primaryKey" json:"email"`
	Name           string    `gorm:"primaryKey" json:"name"`
	Password       string    `gorm:"NOT NULL" json:"-"`
	ProfilePicture string    `json:"profilePicture"`
	BannerPicture  string    `json:"bannerPicture"`
	Introduction   string    `json:"introduction"`
	Exp            uint      `gorm:"NOT NULL; default:0" json:"exp"`
	Management     int       `gorm:"NOT NULL; default:-2" json:"management"` // 當註冊審核中，為-2，成功後會改成0
	SignupIP       string    `gorm:"type:varchar(45)" json:"signupIp"`
	CreatedAt      time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Articles       []Article `gorm:"foreignKey:UserID"`
}

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"NOT NULL" json:"title"`
	Image     string    `json:"image"`
	Content   string    `gorm:"NOT NULL" json:"content"`
	Like      uint      `gorm:"NOT NULL; default:0" json:"like"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags      []Tag     `gorm:"many2many:article_tags"`
	UserID    uint      `gorm:"NOT NULL" json:"userId"`
	User      User
}

type Tag struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"NOT NULL" json:"name"`
	IconName string `gorm:"NOT NULL" json:"iconName"`
}

type Log struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"NOT NULL" json:"type"`
	Message   string    `gorm:"NOT NULL" json:"message"`
	Severity  uint      `gorm:"NOT NULL; default:0" json:"severity"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
}

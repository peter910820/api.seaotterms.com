package gal

import (
	"time"

	"github.com/lib/pq"
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

type Article struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"NOT NULL" json:"title"`
	Image       string         `json:"image"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Content     string         `gorm:"NOT NULL" json:"content"`
	Like        uint           `gorm:"NOT NULL; default:0" json:"like"`
	CreatedAt   time.Time      `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string         `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UpdatedName string         `json:"updatedName"`
}

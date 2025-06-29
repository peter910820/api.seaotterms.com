package blog

import "time"

// all article
type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"NOT NULL" json:"title"`
	Content   string    `gorm:"NOT NULL"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags      []Tag     `gorm:"many2many:article_tags"`
}

// all article tags
type Tag struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"NOT NULL;uniqueIndex" json:"name"`
	IconName string `gorm:"NOT NULL" json:"iconName"`
}

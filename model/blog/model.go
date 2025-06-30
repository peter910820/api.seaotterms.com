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
//不允許修改Name(PK)
type Tag struct {
	Name      string    `gorm:"primaryKey" json:"name"`
	IconName  string    `json:"iconName"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

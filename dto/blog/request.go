package blog

import (
	"time"

	model "api.seaotterms.com/model/blog"
)

type ArticleCreateRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type ArticleUpdateRequest struct {
	Title   string      `gorm:"NOT NULL" json:"title"`
	Content string      `gorm:"NOT NULL" json:"content"`
	Tags    []model.Tag `gorm:"many2many:article_tags" json:"tags"`
}

type TagCreateRequest struct {
	Name     string `json:"name"`
	IconName string `json:"iconName"`
}

type SystemTodoCreateRequest struct {
	SystemName  string     `json:"systemName"`
	Title       string     `json:"title"`
	Detail      string     `json:"detail"`
	Status      uint       `json:"status"`
	Deadline    *time.Time `json:"deadline"`
	Urgency     uint       `json:"urgency"`
	CreatedName string     `json:"createdName"`
}

type SystemTodoUpdateRequest struct {
	SystemName  string     `json:"systemName"`
	Title       string     `json:"title"`
	Detail      string     `json:"detail"`
	Status      uint       `json:"status"`
	Deadline    *time.Time `json:"deadline"`
	Urgency     uint       `json:"urgency"`
	UpdatedName string     `json:"updatedName"`
}

type QuickSystemTodoUpdateRequest struct {
	Status      uint      `json:"status"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UpdatedName string    `json:"updatedName"`
}

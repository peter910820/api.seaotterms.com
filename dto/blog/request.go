package blog

import model "api.seaotterms.com/model/blog"

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

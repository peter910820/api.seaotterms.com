package blog

type ArticleCreateRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type TagCreateRequest struct {
	Name     string `json:"name"`
	IconName string `json:"iconName"`
}

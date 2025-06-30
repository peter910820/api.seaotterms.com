package blog

type ArticleQueryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TagIDs  []uint `json:"tags"`
}

type TagCreateRequest struct {
	Name     string `json:"name"`
	IconName string `json:"iconName"`
}

package blog

type ArticleQueryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TagIDs  []uint `json:"tags"`
}

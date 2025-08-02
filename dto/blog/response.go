package blog

type CommonResponse[T any] struct {
	StatusCode int       `json:"statusCode"` // http status code
	ErrMsg     string    `json:"errMsg"`
	InfoMsg    string    `json:"infoMsg"`
	UserInfo   *UserInfo `json:"userInfo"`
	Data       *T        `json:"data"`
}

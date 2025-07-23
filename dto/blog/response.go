package blog

type CommonResponse[T any] struct {
	StatusCode uint   `json:"statusCode"` // http status code
	ErrMsg     string `json:"errMsg"`
	InfoMsg    string `json:"infoMsg"`
	Data       *T     `json:"data"`
}

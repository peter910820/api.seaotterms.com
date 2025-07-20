package blog

type CommonResponse[T any] struct {
	StatusCode uint   `json:"statusCode"` // http status code
	ErrCode    string `json:"errCode"`
	ErrMsg     string `json:"errMsg"`
	InfoMsg    string `json:"infoMsg"`
	Data       *T     `json:"data"`
}

func CreateDefalutCommonResponse[T any]() CommonResponse[T] {
	return CommonResponse[T]{
		StatusCode: 200,
		ErrCode:    "",
		ErrMsg:     "",
		InfoMsg:    "",
		Data:       nil,
	}
}

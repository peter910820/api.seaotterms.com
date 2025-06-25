package gal

type CommonResponse struct {
	StatusCode uint   `json:"statusCode"` // http status code
	ErrMsg     string `json:"errMsg"`
}

type LoginResponse struct {
	Data CommonResponse
}

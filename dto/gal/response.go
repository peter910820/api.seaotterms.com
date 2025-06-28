package gal

type CommonResponse struct {
	StatusCode uint   `json:"statusCode"` // http status code
	ErrMsg     string `json:"errMsg"`
	InfoMsg    string `json:"infoMsg"`
}

// type LoginResponse struct {
// 	Data CommonResponse
// }

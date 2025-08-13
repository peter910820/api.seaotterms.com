package blog

import "time"

type CommonResponse[T any] struct {
	StatusCode int       `json:"statusCode"` // http status code
	ErrMsg     string    `json:"errMsg"`
	InfoMsg    string    `json:"infoMsg"`
	UserInfo   *UserInfo `json:"userInfo"`
	Data       *T        `json:"data"`
}

type UserQueryResponse struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"NOT NULL unique" json:"username"`
	Exp        int       `gorm:"default:0" json:"exp"`
	Management bool      `gorm:"default:false" json:"management"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdateName string    `json:"updateName"`
}

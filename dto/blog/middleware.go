package blog

import "time"

type UserInfo struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Exp         int       `json:"exp"`
	Management  bool      `json:"management"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdateName  string    `json:"update_name"`
	Avatar      string    `json:"avatar"`
	DataVersion int       `json:"dataVersion"`
}

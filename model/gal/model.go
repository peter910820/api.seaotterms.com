package gal

import (
	"time"
)

// type DownloadArticle struct {
// 	ID           uint      `gorm:"primaryKey" json:"id"`
// 	Image        string    `json:"image"`
// 	Content      string    `gorm:"NOT NULL" json:"content"`
// 	DownloadType string    `gorm:"NOT NULL" json:"downloadType"`
// 	DownloadURL  string    `gorm:"NOT NULL" json:"downloadUrl"`
// 	CreatedAt    time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
// 	CreatedName  string    `gorm:"NOT NULL" json:"createdName"`
// 	UpdatedAt    time.Time `json:"updatedAt"`
// 	UpdatedName  string    `json:"updatedName"`
// }

// galgame brand record schema
type BrandRecord struct {
	Brand       string    `gorm:"primaryKey" json:"brand"`          // PK
	Completed   int       `gorm:"not null" json:"completed"`        // Completed game amount
	Total       int       `gorm:"not null" json:"total"`            // Total game amount
	Annotation  string    `gorm:"not null" json:"annotation"`       // Annotation
	Dissolution bool      `gorm:"default:false" json:"dissolution"` // Dissolution
	InputTime   time.Time `gorm:"autoCreateTime" json:"inputTime"`  // InputTime
	InputName   string    `gorm:"not null" json:"inputName"`        // InputName
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"updateTime"` // UpdateTime
	UpdateName  string    `gorm:"not null" json:"updateName"`       // UpdateName
}

// galgame game record schema
type GameRecord struct {
	Name        string    `gorm:"primaryKey" json:"name"`           // PK
	Brand       string    `gorm:"not null" json:"brand"`            // Brand
	ReleaseDate time.Time `gorm:"not null" json:"releaseDate"`      // ReleaseDate
	AllAges     bool      `gorm:"not null" json:"allAges"`          // For all ages
	EndDate     time.Time `gorm:"not null" json:"endDate"`          // End date of play
	InputTime   time.Time `gorm:"autoCreateTime" json:"inputTime"`  // InputTime
	InputName   string    `gorm:"not null" json:"inputName"`        // InputName
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"updateTime"` // UpdateTime
	UpdateName  string    `gorm:"not null" json:"updateName"`       // UpdateName
}

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Email          string    `gorm:"uniqueIndex" json:"email"`
	UserName       string    `gorm:"uniqueIndex" json:"userName"`
	Password       string    `gorm:"NOT NULL" json:"-"`
	ProfilePicture string    `json:"profilePicture"`
	BannerPicture  string    `json:"bannerPicture"`
	Introduction   string    `json:"introduction"`
	Exp            uint      `gorm:"NOT NULL; default:0" json:"exp"`
	Management     int       `gorm:"NOT NULL; default:-2" json:"management"` // 當註冊審核中，為-2，成功後會改成0
	SignupIP       string    `gorm:"type:varchar(45)" json:"signupIp"`
	CreatedAt      time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Articles       []Article `gorm:"foreignKey:UserID"`
}

type Article struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"NOT NULL" json:"title"`
	Image     string    `json:"image"`
	Content   string    `gorm:"NOT NULL" json:"content"`
	Like      uint      `gorm:"NOT NULL; default:0" json:"like"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Tags      []Tag     `gorm:"many2many:article_tags"`
	UserID    uint      `gorm:"NOT NULL" json:"userId"`
	User      User
}

type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"NOT NULL;uniqueIndex" json:"name"`
	IconName  string    `gorm:"NOT NULL" json:"iconName"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Log struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"NOT NULL" json:"type"`
	Message   string    `gorm:"NOT NULL" json:"message"`
	Severity  uint      `gorm:"NOT NULL; default:0" json:"severity"`
	CreatedAt time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
}

type TmpData struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Type         string    `gorm:"NOT NULL" json:"type"`
	Content      string    `gorm:"NOT NULL" json:"content"`
	ExpirationAt time.Time `gorm:"NOT NULL" json:"expirationAt"`
	CreatedAt    time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
}

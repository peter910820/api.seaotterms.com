package galgame

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

// Galgame Brand資料主表
type Brand struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"NOT NULL" json:"name"`
	WorkAmount  int       `gorm:"NOT NULL; default:0" json:"workAmount"`      // 作品數量
	OfficialUrl string    `json:"officialUrl"`                                // 官網URL
	Dissolution bool      `gorm:"NOT NULL; default:false" json:"dissolution"` // 解散標記
	CreatedAt   time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt   time.Time `json:"updatedAt"` // 建立資料的時候，因為該欄位的名字，所以gorm也會預設填值，無需特別定義NOT NULL
	UpdatedName string    `json:"updatedName"`
}

// Galgame遊戲資料主表
type Game struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"NOT NULL" json:"name"`
	ChineseName     string    `json:"chineseName"`
	BrandID         int       `gorm:"NOT NULL" json:"brandId"`
	AllAges         bool      `gorm:"NOT NULL" json:"allAges"`
	ReleaseDate     time.Time `gorm:"NOT NULL" json:"releaseDate"`
	OpUrl           string    `json:"opUrl"`
	GameDescription string    `json:"gameDescription"`
	VndbEvaluate    *float64  `json:"vndbEvaluate"`
	CreatedAt       time.Time `gorm:"NOT NULL; autoCreateTime" json:"createdAt"`
	CreatedName     string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedName     string    `json:"updatedName"`
}

// 自己的Galgame遊戲評分紀錄主表
type PlayRecord struct {
	ID                   int       `gorm:"primaryKey" json:"id"`
	GameID               int       `gorm:"NOT NULL; uniqueIndex" json:"gameId"`
	EndPlayDate          time.Time `gorm:"NOT NULL" json:"endPlayDate"`
	OpDisplayScore       *float64  `json:"opDisplayScore"`
	OpSongScore          *float64  `json:"opSongScore"`
	OpCompatibilityScore *float64  `json:"opCompatibilityScore"`
	EdDisplayScore       *float64  `json:"edDisplayScore"`
	EdSongScore          *float64  `json:"edSongScore"`
	MusicScore           *float64  `json:"musicScore"`
	PlotScore            float64   `json:"plotScore"`
	ArtScore             float64   `json:"artScore"`
	SystemScore          float64   `json:"systemScore"`
	ThemeScore           float64   `json:"themeScore"`
	ConclusionScore      float64   `json:"conclusionScore"`
	Category             string    `gorm:"NOT NULL; default:一般" json:"category"`
	Recommended          int       `gorm:"NOT NULL; default:0" json:"recommended"`
	CreatedName          string    `gorm:"NOT NULL" json:"createdName"`
	UpdatedName          string    `json:"updatedName"`
}

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

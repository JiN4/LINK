package model

import (
	_ "time"

	"github.com/jinzhu/gorm"
)

// //model内で共通のfieldをここに抜き出す
//type //model struct {
//    ID      uint `gorm:"primary_key" json:"id"`
//    CreatedAt time.Time
//    UpdatedAt time.Time
//}

// 会社情報
type Company struct {
	//    //Model
	gorm.Model
	Thumbnail   string `gorm:"not null"` //nilかもしれない
	Name        string `gorm:"not null"` // 社名
	Description string `gorm:"not null"` // 概要
	Url         string `gorm:"not null"` // HpURL
	Address     string `gorm:"not null"` // 住所
	TopEvent    uint   //nilかもしれない 推したいイベント
	Events      []Event
}

// 開催イベント
type Event struct {
	//   //Model
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Thumbnail   string `gorm:"not null"`
	Icon        string `gorm:"not null"`
	CompanyID   uint  `gorm:"index;not null"`
	Contents    []Content
}

type Events []Event

// コンテンツ
type Content struct {
	//Model
	gorm.Model
	CompanyID    uint `gorm:"index;not null"`
	EventID      uint `gorm:"index;not null"`
	Author       string `gorm:"not null"`
	Title        string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Thumbnail    string // Url
	IsBookMarked bool `gorm:"not null"`
	MovieUrl     string `gorm:"not null"`
	BrowsedNum   uint       `gorm:"default:0"`
	Categories   []Category `gorm:"many2many:contents_categories;"`
}

type ContentsCategories struct {
	gorm.Model
	ContentID  uint `gorm:"index;not null"`
	CategoryID uint `gorm:"index;not null"`
}

type Category struct {
	//Model
	gorm.Model
	//ContentID uint
	Name     string
	Kind     uint
	Contents []Content `gorm:"many2many:contents_categories;"`
}

type Categories []Category

// 投稿記事
type Article struct {
	//Model
	gorm.Model
	ContentID   uint  `gorm:"index;not null"`
    UserID      uint  `gorm:"index;not null"`
	Description string
	Title       string
	Thumbnail   string
	Url         string
}

type Articles []Article

// 自分のユーザー情報
type User struct {
	//Model
	gorm.Model
	UID   string `gorm:"not null;unique"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null"`
	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
}

// マイユーザーと記事の関係テーブル
type MyselfArticles struct {
	//Model
	UserId    uint
	ArticleId uint
	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
}

type BrowsedHistory struct { //
	gorm.Model
	UserId    uint `gorm:"index;not null"`
	ContentId uint `gorm:"index;not null"`
}

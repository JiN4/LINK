package service

import (
	"time"
)

// 開催イベント
type Event struct { //
	EventCard
	Contents ContentCards `json:"contents"`
}

// Event一覧で取得するカード
type EventCard struct {
	ID          uint        `json:"id"`
	Company     CompanyCard `json:"company"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Thumbnail   string      `json:"thumbnail"`
	CompanyID   uint        `json:"-"`
	CreatedAt   time.Time   `json:"created_at"`
}

type EventCards []EventCard

// 会社情報
type Company struct { //
	Id          uint   // 会社id
	thumbnail   string //nilかもしれない
	Name        string // 社名
	Description string // 概要
	Url         string // HpURL
	Address     string // 住所
	// TopEvent    EventCard   //nilかもしれない 推したいイベント
	// Events      []EventCard // イベント一覧
}

type CompanyCard struct {
	Id          uint   `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Address     string `json:"address"`
}

type CompanyCards []CompanyCard

// コンテンツ
type Content struct { //
	ContentCard
	Categories Categories
	// SlideUrls    []string
	MovieUrl string `json:"movie_url"`
	// timeMap      map[int]uint //時間, ページ数
}

type Contents []Content

type ContentCard struct { //
	Id           uint        `json:"id"`
	Company      CompanyCard `json:"company"`
	Author       string      `json:"author"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Thumbnail    string      `json:"thumbnail"`
	IsBookMarked bool        `json:"is_bookmarked"`
	CompanyID    uint        `json:"-"`
	CreatedAt    time.Time   `json:"created_at"`
}

type ContentCards []ContentCard

// 投稿記事
type Article struct { //
	Id          uint   `json:"id"`
	ContentId   uint   `json:"content_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Url         string `json:"url"`
}

type Category struct { //
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Categories []Category

// ホーム画面で欲しいやつ
type HomeContent struct {
	Banners     ContentCards `json:"banners"`
	Events      EventCards   `json:"events"`
	Rankings    ContentCards `json:"rankings"`
	NewContents ContentCards `json:"new_contents"`
	Histories   ContentCards `json:"histories"`
}

// 自分のユーザー情報
type User struct { //
	ID        uint      `json:"id"`
	UID       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Articles  []Article `json:"articles"`
	CreatedAt uint      `created_at`
	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
}

// 閲覧履歴モデル
type BrowsedHistory struct {
	UserID    uint `json:"user_id"`
	ContentID uint `json:"content_id"`
}

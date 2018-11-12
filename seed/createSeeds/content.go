package createSeeds

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/LINK/model"
)

// // コンテンツ
// type Content struct {
// 	//Model
// 	gorm.Model
// 	CompanyId    uint
// 	EventID      uint
// 	Author       string
// 	Title        string
// 	Description  string
// 	Thumbnail    string // Url
// 	IsBookMarked bool
// 	MovieUrl     string
// 	BrowsedNum   uint       `gorm:"default:0"`
// 	Categories   []Category `gorm:"many2many:contents_categories;"`
// }

func CreateSeedContents() {
	company_infos := []map[string]string{
		map[string]string{
			"EventID":      "3",
			"CompanyID":    "1",
			"Title":        "Androidの省電力について考える",
			"Author":       "中西 良明",
			"Description":  "Android 6.0で導入された新しい省電力機能であるDozeモードやApp Standbyについて紹介",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/c58b6a6646054b3d9b8ef814e437659a/slide_0.jpg",
			"IsBookMarked": "true",
			"MovieUrl":     "https://www.youtube.com/watch?v=U7dK-qJ1N9I",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "3",
			"CompanyID":    "1",
			"Title":        "明日から使えるRxJava頻出パターン",
			"Author":       "kazy",
			"Description":  "頻出設計のパターンの紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigirxjava-160218013345/95/rxjava-droid-kaigi-2016-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=_Rp6kNxAHrU",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "3",
			"CompanyID":    "1",
			"Title":        "TV向けAndroidアプリの開発TIPSと最新事情",
			"Author":       "小林 剛士",
			"Description":  "既にお持ちのアプリをTV向けに開発すべき手段や機能などをご紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigi2016-tvandroid-160222035428/95/droidkaigi2016-tvandroidtips-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=lpjFL6o-Ltw",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "3",
			"CompanyID":    "1",
			"Title":        "Go MobileでAndroidアプリ開発",
			"Author":       "tenntenn",
			"Description":  "Go Mobileのソースコードやドキュメント、海外のカンファレンスの発表資料などから調べた内容についてまとめ、何ができるのか、どうやって実装するのかについてGo初心者でも分かるように説明します",
			"Thumbnail":    "https://image.slidesharecdn.com/go-mobileandroid-160217232100/95/go-mobileandroid-1-638.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=V7bwZQ6S4VE",
			"Categories":   "Go,Android",
		},
		map[string]string{
			"EventID":      "3",
			"CompanyID":    "1",
			"Title":        "僕がテスト書け書けおじさんになった経緯とその過程でやったこと",
			"Author":       "yuyakaido",
			"Description":  "とあるプロジェクトでの例をもとに、テストコードがまったくなく気合でメンテナンスしている状態からどのようにテスタビリティの高い設計に変更し、どのようなテストフレームワークを導入し、どのようにチーム全員でテストを書く文化を作ったのかということを話したいと思います。",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/2f64526f0b364719b73c37513aaacecc/slide_0.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=dP1beWS2148",
			"Categories":   "Java,Android",
		},

		map[string]string{
			"EventID":      "1",
			"CompanyID":    "1",
			"Title":        "Androidの省電力について考える",
			"Author":       "中西 良明",
			"Description":  "Android 6.0で導入された新しい省電力機能であるDozeモードやApp Standbyについて紹介",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/c58b6a6646054b3d9b8ef814e437659a/slide_0.jpg",
			"IsBookMarked": "true",
			"MovieUrl":     "https://www.youtube.com/watch?v=U7dK-qJ1N9I",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "1",
			"CompanyID":    "1",
			"Title":        "明日から使えるRxJava頻出パターン",
			"Author":       "kazy",
			"Description":  "頻出設計のパターンの紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigirxjava-160218013345/95/rxjava-droid-kaigi-2016-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=_Rp6kNxAHrU",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "1",
			"CompanyID":    "1",
			"Title":        "TV向けAndroidアプリの開発TIPSと最新事情",
			"Author":       "小林 剛士",
			"Description":  "既にお持ちのアプリをTV向けに開発すべき手段や機能などをご紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigi2016-tvandroid-160222035428/95/droidkaigi2016-tvandroidtips-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=lpjFL6o-Ltw",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "1",
			"CompanyID":    "1",
			"Title":        "Go MobileでAndroidアプリ開発",
			"Author":       "tenntenn",
			"Description":  "Go Mobileのソースコードやドキュメント、海外のカンファレンスの発表資料などから調べた内容についてまとめ、何ができるのか、どうやって実装するのかについてGo初心者でも分かるように説明します",
			"Thumbnail":    "https://image.slidesharecdn.com/go-mobileandroid-160217232100/95/go-mobileandroid-1-638.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=V7bwZQ6S4VE",
			"Categories":   "Go,Android",
		},
		map[string]string{
			"EventID":      "1",
			"CompanyID":    "1",
			"Title":        "僕がテスト書け書けおじさんになった経緯とその過程でやったこと",
			"Author":       "yuyakaido",
			"Description":  "とあるプロジェクトでの例をもとに、テストコードがまったくなく気合でメンテナンスしている状態からどのようにテスタビリティの高い設計に変更し、どのようなテストフレームワークを導入し、どのようにチーム全員でテストを書く文化を作ったのかということを話したいと思います。",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/2f64526f0b364719b73c37513aaacecc/slide_0.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=dP1beWS2148",
			"Categories":   "Java,Android",
		},

		map[string]string{
			"EventID":      "2",
			"CompanyID":    "1",
			"Title":        "Androidの省電力について考える",
			"Author":       "中西 良明",
			"Description":  "Android 6.0で導入された新しい省電力機能であるDozeモードやApp Standbyについて紹介",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/c58b6a6646054b3d9b8ef814e437659a/slide_0.jpg",
			"IsBookMarked": "true",
			"MovieUrl":     "https://www.youtube.com/watch?v=U7dK-qJ1N9I",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "2",
			"CompanyID":    "1",
			"Title":        "明日から使えるRxJava頻出パターン",
			"Author":       "kazy",
			"Description":  "頻出設計のパターンの紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigirxjava-160218013345/95/rxjava-droid-kaigi-2016-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=_Rp6kNxAHrU",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "2",
			"CompanyID":    "1",
			"Title":        "TV向けAndroidアプリの開発TIPSと最新事情",
			"Author":       "小林 剛士",
			"Description":  "既にお持ちのアプリをTV向けに開発すべき手段や機能などをご紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigi2016-tvandroid-160222035428/95/droidkaigi2016-tvandroidtips-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=lpjFL6o-Ltw",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "2",
			"CompanyID":    "1",
			"Title":        "Go MobileでAndroidアプリ開発",
			"Author":       "tenntenn",
			"Description":  "Go Mobileのソースコードやドキュメント、海外のカンファレンスの発表資料などから調べた内容についてまとめ、何ができるのか、どうやって実装するのかについてGo初心者でも分かるように説明します",
			"Thumbnail":    "https://image.slidesharecdn.com/go-mobileandroid-160217232100/95/go-mobileandroid-1-638.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=V7bwZQ6S4VE",
			"Categories":   "Go,Android",
		},
		map[string]string{
			"EventID":      "2",
			"CompanyID":    "1",
			"Title":        "僕がテスト書け書けおじさんになった経緯とその過程でやったこと",
			"Author":       "yuyakaido",
			"Description":  "とあるプロジェクトでの例をもとに、テストコードがまったくなく気合でメンテナンスしている状態からどのようにテスタビリティの高い設計に変更し、どのようなテストフレームワークを導入し、どのようにチーム全員でテストを書く文化を作ったのかということを話したいと思います。",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/2f64526f0b364719b73c37513aaacecc/slide_0.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=dP1beWS2148",
			"Categories":   "Java,Android",
		},

		map[string]string{
			"EventID":      "4",
			"CompanyID":    "1",
			"Title":        "Androidの省電力について考える",
			"Author":       "中西 良明",
			"Description":  "Android 6.0で導入された新しい省電力機能であるDozeモードやApp Standbyについて紹介",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/c58b6a6646054b3d9b8ef814e437659a/slide_0.jpg",
			"IsBookMarked": "true",
			"MovieUrl":     "https://www.youtube.com/watch?v=U7dK-qJ1N9I",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "4",
			"CompanyID":    "1",
			"Title":        "明日から使えるRxJava頻出パターン",
			"Author":       "kazy",
			"Description":  "頻出設計のパターンの紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigirxjava-160218013345/95/rxjava-droid-kaigi-2016-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=_Rp6kNxAHrU",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "4",
			"CompanyID":    "1",
			"Title":        "TV向けAndroidアプリの開発TIPSと最新事情",
			"Author":       "小林 剛士",
			"Description":  "既にお持ちのアプリをTV向けに開発すべき手段や機能などをご紹介",
			"Thumbnail":    "https://image.slidesharecdn.com/droidkaigi2016-tvandroid-160222035428/95/droidkaigi2016-tvandroidtips-1-638.jpg",
			"IsBookMarked": "false",
			"MovieUrl":     "https://www.youtube.com/watch?v=lpjFL6o-Ltw",
			"Categories":   "Java,Android",
		},
		map[string]string{
			"EventID":      "4",
			"CompanyID":    "1",
			"Title":        "Go MobileでAndroidアプリ開発",
			"Author":       "tenntenn",
			"Description":  "Go Mobileのソースコードやドキュメント、海外のカンファレンスの発表資料などから調べた内容についてまとめ、何ができるのか、どうやって実装するのかについてGo初心者でも分かるように説明します",
			"Thumbnail":    "https://image.slidesharecdn.com/go-mobileandroid-160217232100/95/go-mobileandroid-1-638.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=V7bwZQ6S4VE",
			"Categories":   "Go,Android",
		},
		map[string]string{
			"EventID":      "4",
			"CompanyID":    "1",
			"Title":        "僕がテスト書け書けおじさんになった経緯とその過程でやったこと",
			"Author":       "yuyakaido",
			"Description":  "とあるプロジェクトでの例をもとに、テストコードがまったくなく気合でメンテナンスしている状態からどのようにテスタビリティの高い設計に変更し、どのようなテストフレームワークを導入し、どのようにチーム全員でテストを書く文化を作ったのかということを話したいと思います。",
			"Thumbnail":    "https://speakerd.s3.amazonaws.com/presentations/2f64526f0b364719b73c37513aaacecc/slide_0.jpg",
			"IsBookMarked": "TRUE",
			"MovieUrl":     "https://www.youtube.com/watch?v=dP1beWS2148",
			"Categories":   "Java,Android",
		},
	}

	for _, info := range company_infos {
		eventID, _ := strconv.Atoi(info["EventID"])
		companyID, _ := strconv.Atoi(info["CompanyID"])
		isBookMarked, _ := strconv.ParseBool(info["IsBookMarked"])
		categories, _ := model.GetCategoriesByNameList(strings.Split(info["Categories"], ","))
		createContent(model.Content{
			EventID:      uint(eventID),
			CompanyID:    uint(companyID),
			Title:        info["Title"],
			Author:       info["Author"],
			Description:  info["Description"],
			Thumbnail:    info["Thumbnail"],
			IsBookMarked: isBookMarked,
			MovieUrl:     info["MovieUrl"],
			BrowsedNum:   uint(rand.Intn(1000)),
			Categories:   categories,
		})
	}
}

func createContent(c model.Content) {
	content, err := model.CreateContent(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content created: %v\n", content.CreatedAt)
}

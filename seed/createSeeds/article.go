package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/LINK/model"
)

// // 投稿記事
// type Article struct {
// 	//Model
// 	ContentId   uint
// 	Description string
// 	Title       string
// 	Thumbnail   string
// 	Url         string
// }

func CreateSeedArticle() {
	article_infos := []map[string]string{
		map[string]string{
			"ContentId":   "1",
			"UserID":      "1",
			"Title":       "Androidアプリを超省電力で作ってみた",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://i2.wp.com/qs.nndo.jp/wp-content/uploads/2017/10/qiita-fb-2887e7b4aad86fd8c25cea84846f2236.png?fit=200%2C200",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "1",
			"Title":       "Androidアプリを超省電力で作ってみる",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://akaneiro.me/wp-content/uploads/2016/01/notec.png",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "1",
			"Title":       "Androidアプリを超省電力で作ってみよう",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://i2.wp.com/qs.nndo.jp/wp-content/uploads/2017/10/qiita-fb-2887e7b4aad86fd8c25cea84846f2236.png?fit=200%2C200",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "1",
			"Title":       "Androidアプリを超省電力で作ってみたい",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://akaneiro.me/wp-content/uploads/2016/01/notec.png",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "2",
			"Title":       "Androidアプリを超省電力で作ってみた気がした",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://i2.wp.com/qs.nndo.jp/wp-content/uploads/2017/10/qiita-fb-2887e7b4aad86fd8c25cea84846f2236.png?fit=200%2C200",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "2",
			"Title":       "Androidアプリを超省電力で作ってみたね",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://akaneiro.me/wp-content/uploads/2016/01/notec.png",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "3",
			"Title":       "Androidアプリを超省電力で作ってみたような",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://akaneiro.me/wp-content/uploads/2016/01/notec.png",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "3",
			"Title":       "Androidアプリを超省電力で作ってみた！",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://i2.wp.com/qs.nndo.jp/wp-content/uploads/2017/10/qiita-fb-2887e7b4aad86fd8c25cea84846f2236.png?fit=200%2C200",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "3",
			"Title":       "Androidアプリを超省電力で作ってみた！！",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://i2.wp.com/qs.nndo.jp/wp-content/uploads/2017/10/qiita-fb-2887e7b4aad86fd8c25cea84846f2236.png?fit=200%2C200",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
		map[string]string{
			"ContentId":   "1",
			"UserID":      "3",
			"Title":       "Androidアプリを超省電力で作ってみた？",
			"Description": "DroidKaigi2016で得た知識を生かしAndroidアプリを超省電力で作成",
			"Thumbnail":   "https://akaneiro.me/wp-content/uploads/2016/01/notec.png",
			"Url":         "https://qiita.com/joji/items/57113f631f90f1e4e458",
		},
	}

	for _, info := range article_infos {
		contentId, _ := strconv.Atoi(info["ContentId"])

		userId, _ := strconv.Atoi(info["UserID"])
		createArticle(model.Article{
			ContentID:   uint(contentId),
			UserID:      uint(userId),
			Title:       info["Title"],
			Description: info["Description"],
			Thumbnail:   info["Thumbnail"],
			Url:         info["Url"],
		})
	}
}

func createArticle(article model.Article) {
	article, err := model.CreateArticle(article)
	if err != nil {
		panic(err)
	}
	fmt.Printf("article created\n")
}

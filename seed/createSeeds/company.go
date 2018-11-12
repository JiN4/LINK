package createSeeds

import (
	"fmt"

	"github.com/LINK/model"
)

// type Company struct { //
// 	Id          uint        // 会社id
// 	thumbnail   string      //nilかもしれない
// 	Name        string      // 社名
// 	Description string      // 概要
// 	Url         string      // HpURL
// 	Address     string      // 住所
// 	TopEvent    EventCard   //nilかもしれない 推したいイベント
// 	Events      []EventCard // イベント一覧
// }

func CreateSeedCompanies() {
	company_infos := []map[string]string{
		map[string]string{
			"Name":        "DroidKaigi",
			"Thumbnail":   "",
			"Description": "エンジニアが主役のAndroidカンファレンス",
			"Url":         "droidkaigi.github.io",
			"Address":     "東京都新宿区西新宿 8−17−3 住友不動産新宿グランドタワー",
		},
		map[string]string{
			"Name":        "Go Conference",
			"Thumbnail":   "",
			"Description": "プログラミング言語Goの日本最大のカンファレンス",
			"Url":         "https://gocon.connpass.com/",
			"Address":     "東京都港区六本木6-10-1",
		},
		map[string]string{
			"Name":        "GTC Japan",
			"Thumbnail":   "",
			"Description": "日本最大のGPUテクノロジーイベント",
			"Url":         "https://www.nvidia.com/ja-jp/gtc/",
			"Address":     "東京都港区高輪3-13-1グランドプリンスホテル新高輪",
		},
		map[string]string{
			"Name":        "一般社団法人PyCon JP",
			"Thumbnail":   "",
			"Description": "Pythonに関する国際カンファレンス",
			"Url":         "https://www.pycon.jp/",
			"Address":     "東京都台東区東上野1-11-1 GOSHO春日通りビル9F",
		},
	}

	for _, info := range company_infos {
		//kindId, _ := strconv.Atoi(info["Kind"])
		createCompany(model.Company{
			Name:        info["Name"],
			Thumbnail:   info["Thumbnail"],
			Description: info["Description"],
			Url:         info["Url"],
			Address:     info["Address"],
		})
	}
}

func createCompany(c model.Company) {
	company, err := model.CreateCompany(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("company created: %v\n", company.CreatedAt)
}

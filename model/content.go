package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/LINK/service"
)

// Create コンテントを作成する
func CreateContent(content Content) (Content, error) {
	err := db.Create(&content).Error
	if err != nil {
		return Content{}, err
	}
	return content, nil
}

// sortされたイベントを取得する
func GetSortedContentCards(sort_type string) (service.ContentCards, error) {
	order_type := "created_at"
	switch sort_type {
	case "browsed_num":
		{
			log.Println("Im on Browsed Num")
			order_type = "browsed_num"
		}
	case "new":
		{
			log.Println("Im on New")
			order_type = "created_at"
		}
	default:
		{ // history TODO とりあえず適当に返す
			log.Println("Im on Default")
		}
	}

	contentsOnCompany := ContentsOnCompany{}
	// クエリ直書き作戦でN+1問題を回避
	sql := "SELECT cn.id, cn.author, cn.title, cn.description, cn.thumbnail, cn.is_book_marked, cn.created_at, cm.id AS company_id, cm.thumbnail AS company_thumbnail, cm.name, cm.description company_description, cm.url, cm.address FROM contents AS cn INNER JOIN companies AS cm ON cn.company_id = cm.id ORDER  BY " + order_type + " DESC"
	err := db.Raw(sql).Scan(&contentsOnCompany).Error

	return contentsOnCompany.Cards(), err
}

// ContentCardsをEventIdでフィルターかけて返す
func GetContentCardsFilteredByEventId(eventId uint) (service.ContentCards, error) {
	log.Println("onGetContentCardsFilteredByEventId")
	log.Println("eventId: ", eventId)
	contentCards := service.ContentCards{}
	err := db.Table("contents").Where("event_id = ?", eventId).Find(&contentCards).Error
	return contentCards, err
}

func GetContentsFilteredByCategory(name string) ([]Content, error) {
	contents := []Content{}
	category, _ := GetCategoryByName(name)
	log.Println("Category", category.ID)
	// クエリ直書き作戦でN+1問題を回避
	sql := "SELECT contents.id, author, title, description, thumbnail, is_book_marked, movie_url, browsed_num, company_id FROM contents LEFT JOIN contents_categories ON contents_categories.content_id = contents.id LEFT JOIN categories AS cate ON contents_categories.category_id = cate.id WHERE (cate.id = 2 OR  cate.id = 3) AND (cate.id = 4)"
	log.Println("sql: " + sql)
	db.Raw(sql).Scan(&contents)
	// db.Raw("SELECT * FROM contents INNER JOIN contents_categories ON contents_categories.content_id = contents.id LEFT JOIN categories ON contents_categories.category_id = categories.id").Scan(&contents)

	log.Println("CategorizedContents", contents)
	return contents, nil
}

func GetAllContentCards() (service.ContentCards, error) {
	// contentCards := service.ContentCards{}
	// err := db.Table("contents").Find(&contentCards).Error
	// for i, contentCard := range contentCards {
	// 	contentCards[i].Company, _ = GetCompanyCardById(contentCard.CompanyID)
	// }
	// return contentCards, err
	contentsOnCompany := ContentsOnCompany{}
	// クエリ直書き作戦でN+1問題を回避
	sql := "SELECT cn.id, cn.author, cn.title, cn.description, cn.thumbnail, cn.is_book_marked, cn.created_at, cm.id AS company_id, cm.thumbnail AS company_thumbnail, cm.name, cm.description company_description, cm.url, cm.address FROM contents AS cn INNER JOIN companies AS cm ON cn.company_id = cm.id"
	err := db.Raw(sql).Scan(&contentsOnCompany).Error

	return contentsOnCompany.Cards(), err
}

func GetEventCardsByBrowsedHistory(userID string) (service.ContentCards, error) {
	contentsOnCompany := ContentsOnCompany{}
	// クエリ直書き作戦でN+1問題を回避
	sql := "SELECT DISTINCT * FROM contents INNER JOIN browsed_histories ON browsed_histories.content_id = contents.id WHERE browsed_histories.user_id = " + userID + " ORDER BY browsed_histories.created_at DESC"
	err := db.Raw(sql).Scan(&contentsOnCompany).Error

	return contentsOnCompany.Cards(), err
}


// CategoryIdのListでFilterしたContentsListを返す
func GetContentsFilteredByCategoryIdList(platformCategoryIds []string, languageCategoryIds []string) (service.ContentCards, error) {
	contentCards := service.ContentCards {}
	platformSubQuery := getSubQueryForFilterByCategoryId(platformCategoryIds)
	languageSubQuery := getSubQueryForFilterByCategoryId(languageCategoryIds)

	subSql := ""
	if platformSubQuery != "" && languageSubQuery != "" {
		subSql = "WHERE (" + platformSubQuery + " AND " + languageSubQuery + ")"
	} else if platformSubQuery == "" && languageSubQuery != "" {
		subSql = "WHERE " + languageSubQuery
	} else if platformSubQuery != ""  && languageSubQuery == "" {
		subSql = "WHERE " + platformSubQuery
	} else {
		subSql = ""
	}

	// クエリ直書き作戦でN+1問題を回避
	sql := "SELECT contents.id, author, title, description, thumbnail, is_book_marked, movie_url, browsed_num FROM contents LEFT JOIN contents_categories ON contents_categories.content_id = contents.id LEFT JOIN categories AS cate ON contents_categories.category_id = cate.id " + subSql
	fmt.Println("sql: " + sql)
	db.Raw(sql).Scan(&contentCards)

	log.Println("CategorizedContents", contentCards)
	return contentCards, nil
}


// WHERE (cate.id = 2 OR  cate.id = 3) AND (cate.id = 4)
func getSubQueryForFilterByCategoryId(ids []string) (string) {
	if(ids[0] != "") {
		subsql := "( "
			for _, id := range ids {
				subsql += "cate.id = " + id + " OR "
			}
			subsql = strings.TrimRight(subsql, "OR ") + ")"

			return subsql
	} else {
			return ""
	}
}


// ContentをIDで指定して取得する
func GetContentById(contentId uint) (service.Content, error) {
	content := service.Content{}
	err := db.Find(&content, contentId).Error
	companyCard, _ := GetCompanyCardById(content.CompanyID)
	content.Company = companyCard
	categories, _ := GetCategoriesByContentId(content.Id)
	content.Categories = categories
	return content, err
}

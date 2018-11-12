package model

import (
    "log"
)

// Create articleを作成する
func CreateArticle(article Article) (Article, error) {
	err := db.Create(&article).Error
	if err != nil {
		return Article{}, err
	}
	return article, nil
}

// Articleをcontent_idで指定して取得する
func GetArticlesByContentId(contentId uint) ([]Article, error) {
	articles := []Article{}
	err := db.Where("content_id = ?", contentId).Find(&articles).Error
	return articles, err

}

// Articleをcontent_idで指定して取得する
func GetArticlesByUID(uid string) (Articles, error) {
    user := User{}
    err := db.Where("uid = ?", uid).Find(&user).Error
    if err != nil {
        log.Println("ERROR: ", err)
        return Articles{}, err
    }
	articles := []Article{}
	//err = db.Where("user_id = ?", user.ID).Find(&articles).Error
	err = db.Where("user_id = ?", 0).Find(&articles).Error
	return articles, err
}

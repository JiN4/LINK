package model

import (
	"log"

	"github.com/LINK/service"
)

// Create categoryを作成する
func CreateCategory(category Category) (Category, error) {
	err := db.Create(&category).Error
	if err != nil {
		return Category{}, err
	}
	return category, nil
}

func GetCategoriesByContentId(contentId uint) (service.Categories, error) {
	categories := service.Categories{}
	sql := "SELECT id, name FROM categories LEFT JOIN contents_categories ON contents_categories.category_id = categories.id WHERE contents_categories.content_id = 2"
	log.Println("カテゴリ検索:\n", sql)
	err := db.Raw(sql).Scan(&categories).Error

	return categories, err
}

func GetAllCategories() ([]Category, error) {
	categories := make([]Category, 0)
	err := db.Find(&categories).Error
	return categories, err
}

func GetCategoryByName(name string) (Category, error) {
	category := Category{}
	err := db.Where("name = ?", name).First(&category).Error
	return category, err
}

func GetCategoriesByNameList(names []string) ([]Category, error) {
	categories := []Category{}
	for _, name := range names {
		category, _ := GetCategoryByName(name)
		categories = append(categories, category)
	}

	return categories, nil
}

// PlatFormのカテゴリを一覧取得
func GetPlatformCategories() service.Categories {
	return getServiceCategoriesFilteredByKind(0)
}

// KindがLanguageのカテゴリを一覧取得
func GetLanguageCategories() service.Categories {
	return getServiceCategoriesFilteredByKind(1)
}

// ***************** Private Methods ****************** //
// カテゴリーをKindでフィルターしたものを返す
func getServiceCategoriesFilteredByKind(kindId uint) service.Categories {
	log.Println("kindId = ", kindId)
	categories := service.Categories{}
	err := db.Where("kind = ?", kindId).Find(&categories).Error
	if err != nil {
		log.Println("Errored On getCategoriesFilteredByKind")
		log.Println(err)
		return service.Categories{}
	}
	return categories
}

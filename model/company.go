package model

import (
	"github.com/LINK/service"
)

// Create companyを作成する
func CreateCompany(company Company) (Company, error) {
	err := db.Create(&company).Error
	if err != nil {
		return Company{}, err
	}
	return company, nil
}

// CompanyをIDで指定して取得する
func GetCompanyById(companyId uint) (Company, error) {
	company := Company{}
	err := db.Find(&company, companyId).Error
	return company, err
}

func GetServiceCompanyCardById(id uint) (service.CompanyCard, error) {
	companyCard := service.CompanyCard{}
	err := db.Table("companies").Find(&companyCard, id).Error

	return companyCard, err
}

// CompanyCardIdで検索してを取得する
func GetCompanyCardById(id uint) (service.CompanyCard, error) {
	companyCard := service.CompanyCard{}
	err := db.Table("companies").Find(&companyCard, id).Error

	return companyCard, err
}

func GetAllCompanyCards() (service.CompanyCards, error) {
	companyCards := service.CompanyCards{}
	err := db.Table("companies").Find(&companyCards).Error

	return companyCards, err
}

func (company Company) Card() service.CompanyCard {
	var card service.CompanyCard
	card = service.CompanyCard{
		Id:          company.ID,
		Thumbnail:   company.Thumbnail,
		Name:        company.Name,
		Description: company.Description,
		Url:         company.Url,
		Address:     company.Address,
	}
	return card
}

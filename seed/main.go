package main

import (
	"github.com/LINK/seed/createSeeds"
)

func main() {
	createSeeds.CreateSeedCategories()
	createSeeds.CreateSeedCompanies()
	createSeeds.CreateSeedEvents()
	createSeeds.CreateSeedContents()
	createSeeds.CreateSeedArticle()
	createSeeds.CreateSeedUsers()

}

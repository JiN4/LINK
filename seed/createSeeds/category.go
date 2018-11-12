package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/LINK/model"
)

func CreateSeedCategories() {
	category_infos := []map[string]string{
		map[string]string{
			"Name": "C",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Java",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Ruby",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Python",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Go",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Kotlin",
			"Kind": "1",
		},
		map[string]string{
			"Name": "JavaScript",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Scala",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Php",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Swift",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Rust",
			"Kind": "1",
		},
		map[string]string{
			"Name": "Web",
			"Kind": "0",
		},
		map[string]string{
			"Name": "iOS",
			"Kind": "0",
		},
		map[string]string{
			"Name": "Android",
			"Kind": "0",
		},
		map[string]string{
			"Name": "WindowsPhone",
			"Kind": "0",
		},
	}

	for _, info := range category_infos {
		kindId, _ := strconv.Atoi(info["Kind"])
		createCategory(model.Category{
			Name: info["Name"],
			Kind: uint(kindId),
		})
	}
}

func createCategory(c model.Category) {
	category, err := model.CreateCategory(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("category created: %v\n", category.CreatedAt)
}

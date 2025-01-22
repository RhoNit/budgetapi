package main

import (
	"fmt"

	"github.com/RhoNit/budgetapi/cmd/api/requests"
	"github.com/RhoNit/budgetapi/cmd/api/services"
	"github.com/RhoNit/budgetapi/common"
)

func main() {
	db, err := common.NewPostgres()
	if err != nil {
		return
	}

	catSvc := services.NewCategoryService(db)

	categories := []string{
		"Food", "Fruits", "Gym & Sports", "Skin & Hair Care", "Tech Magazine",
		"House Rent", "Electricity Bill", "Transportation", "Clothing", "Travel & Vacations",
		"OTT Subscription", "Emergency Fund", "Insurance", "Investments", "Savings",
	}

	for _, category := range categories {
		_, err = catSvc.CreateCategory(&requests.CategoryRequest{
			Name:     category,
			IsCustom: false,
		})

		if err != nil {
			panic(err)
		}

		fmt.Println("Category: " + category + " created")
	}
}

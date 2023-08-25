package product_service

import (
	"bitbucket.org/oneiota/platform-technical-task/model"
	"bitbucket.org/oneiota/platform-technical-task/repo"
)

func GetProducts() []*model.Product {
	return product_repo.ReadProducts()
}

func GetProduct(PLU string) *model.Product {
	// Get product from CSV based on a given id.

	products := product_repo.ReadProducts()

	for _, product := range products {
		if product.PLU == PLU {
			return product
		}
	}
	return nil
}

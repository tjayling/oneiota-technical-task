package product_repo

import (
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"bitbucket.org/oneiota/platform-technical-task/model"
	size_order_util "bitbucket.org/oneiota/platform-technical-task/util"
)

func ReadProducts() []*model.Product {
	fileName := "products.csv"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	// defer keyword used to ensure the file is closed once the rest of the function is finished processing
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	products := []*model.Product{} // Make empty array for products to be appended to

	currentRecord := getRecord(reader)
	previousRecord := currentRecord

	sizes := []model.Sizes{}

	for i := 0; ; i++ {
		previousPLU := previousRecord[1]
		currentPLU := currentRecord[1]

		if previousPLU == currentPLU {
			sizes = append(sizes, getSize(currentRecord))
		} else {
			product := createProduct(sizes, previousRecord, previousPLU)
			products = append(products, product)
			sizes = []model.Sizes{getSize(currentRecord)}
		}

		previousRecord = currentRecord
		currentRecord = getRecord(reader)
		if currentRecord == nil {
			product := createProduct(sizes, previousRecord, previousPLU)
			products = append(products, product)
			break
		}
	}

	return products
}

func createProduct(sizes model.SizeSlice, previousRecord []string, previousPLU string) *model.Product {
	sortSizes(sizes, previousRecord[4])
	return &model.Product{
		PLU:   previousPLU,
		Name:  previousRecord[2],
		Sizes: sizes,
	}
}

func sortSizes(sizes []model.Sizes, sizeSort string) {
	order := map[string]int{}
	switch sizeSort {
	case "SHOE_UK":
		order = size_order_util.GetEuShoeSizeOrder(sizes)
	case "SHOE_EU":
		order = size_order_util.GetEuShoeSizeOrder(sizes)
	case "CLOTHING_SHORT":
		order = size_order_util.GetClothingOrder()
	}

	sort.Sort(model.CustomSizeSort{Sizes: sizes, Order: order})
}

func getSize(record []string) model.Sizes {
	SKU, err := strconv.Atoi(record[0])
	if err != nil {
		log.Println(err)
	}

	return model.Sizes{
		SKU:  SKU,
		Size: record[3],
	}
}

func getRecord(reader *csv.Reader) []string {
	record, err := reader.Read()
	if err != nil {
		return nil
	}
	for i, field := range record {
		cleaned := strings.TrimSpace(field)
		cleaned = strings.Trim(cleaned, "\"")

		record[i] = cleaned
	}
	return record
}

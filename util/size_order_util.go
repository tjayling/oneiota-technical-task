package size_order_util

import (
	"sort"

	"bitbucket.org/oneiota/platform-technical-task/model"
)

func GetEuShoeSizeOrder(sizes []model.Sizes) map[string]int {
	order := make(map[string]int)
	sortedSizes := make([]model.Sizes, len(sizes))
	copy(sortedSizes, sizes)
	sort.Sort(model.SizeSlice(sortedSizes))

	for i, size := range sortedSizes {
		order[size.Size] = i
	}

	return order
}

// func GetUkShoeSizeOrder(sizes []model.Sizes) map[string]int {
// 	// return map[string]int{
// 	// 	"1 (Child)":    0,
// 	// 	"1.5 (Child)":  0,
// 	// 	"2 (Child)":    0,
// 	// 	"2.5 (Child)":  0,
// 	// 	"3 (Child)":    0,
// 	// 	"3.5 (Child)":  0,
// 	// 	"4 (Child)":    0,
// 	// 	"4.5 (Child)":  0,
// 	// 	"5 (Child)":    0,
// 	// 	"6.5 (Child)":  0,
// 	// 	"7 (Child)":    0,
// 	// 	"7.5 (Child)":  0,
// 	// 	"8 (Child)":    0,
// 	// 	"8.5 (Child)":  0,
// 	// 	"9 (Child)":    0,
// 	// 	"9.5 (Child)":  0,
// 	// 	"10 (Child)":   0,
// 	// 	"10.5 (Child)": 0,
// 	// 	"11 (Child)":   0,
// 	// 	"11.5 (Child)": 0,
// 	// 	"12 (Child)":   0,
// 	// }
// }

func GetClothingOrder() map[string]int {
	return map[string]int{
		"XS":    0,
		"S":     1,
		"M":     2,
		"L":     3,
		"XL":    4,
		"XXL":   5,
		"XXXL":  6,
		"XXXXL": 7,
	}
}

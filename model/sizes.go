package model

import (
	"strconv"
	"strings"
)

type Sizes struct {
	SKU  int    `json:"SKU"`
	Size string `json:"size"`
}

type SizeSlice []Sizes

func (sa SizeSlice) Len() int {
	return len(sa)
}

func (sa SizeSlice) Less(i, j int) bool {
	// Add some logic to see if the size is a child size
	isIChildSize := strings.Contains(sa[i].Size, "(Child)")
	isJChildSize := strings.Contains(sa[j].Size, "(Child)")

	if isIChildSize && !isJChildSize {
		return true
	} else if !isIChildSize && isJChildSize {
		return false
	}

	valueI, errI := strconv.ParseFloat(strings.Split(sa[i].Size, " ")[0], 64)
	valueJ, errJ := strconv.ParseFloat(strings.Split(sa[j].Size, " ")[0], 64)

	if errI == nil && errJ == nil {
		return valueI < valueJ
	}

	return sa[i].Size < sa[j].Size
}

func (sa SizeSlice) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

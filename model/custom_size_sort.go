package model

type CustomSizeSort struct {
	Sizes []Sizes
	Order map[string]int
}

func (css CustomSizeSort) Len() int {
	return len(css.Sizes)
}

func (css CustomSizeSort) Less(i, j int) bool {
	return css.Order[css.Sizes[i].Size] < css.Order[css.Sizes[j].Size]
}

func (css CustomSizeSort) Swap(i, j int) {
	css.Sizes[i], css.Sizes[j] = css.Sizes[j], css.Sizes[i]
}

package model

type Product struct {
	PLU   string  `json:"PLU"`
	Name  string  `json:"name"`
	Sizes []Sizes `json:"sizes"`
}

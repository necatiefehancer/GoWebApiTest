package models

type Product struct {
	ProductName           string  `json:"ProductName"`
	ProductCode           string  `json:"ProductCode"`
	ProductId             string  `json:"ProductId"`
	ProductBuyPrice       float64 `json:"ProductBuyPrice"`
	ProductSellPrice      float64 `json:"ProductSellPrice"`
	ProductProductionFirm string  `json:"ProductProductionFirm"`
	ProductDescription    string  `json:"ProductDescription"`
}

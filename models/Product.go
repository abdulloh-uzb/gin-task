package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category int     `json:"category"`
	Type     int     `json:"type"`
	Price    float64 `json:"price"`
	Stores   []Store `json:"stores"`
}

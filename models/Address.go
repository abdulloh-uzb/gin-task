package models

type Address struct {
	ID       int    `json:"id"`
	District string `json:"district"`
	Street   string `json:"street"`
	StoreID  int    `json:"store_id"`
}

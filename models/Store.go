package models

type Store struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Addresses []Address `json:"addresses"`
}

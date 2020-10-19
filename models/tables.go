package models

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Stock uint   `json:"stock"`
}

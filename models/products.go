package models

type Product struct {
	Id          int   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
package model

type Book struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Data        string  `json:"data"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

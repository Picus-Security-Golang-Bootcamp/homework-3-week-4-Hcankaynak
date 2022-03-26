package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PageNumber int    `json:"pageNumber"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	StockCode  string `json:"stockCode"`
	ISBN       string `json:"isbn"`
	AuthorId   int    `json:"authorId"`
}

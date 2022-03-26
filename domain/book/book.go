package book

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) (*BookRepository, error) {
	bookRepo := BookRepository{db: db}
	bookRepo.db.AutoMigrate(&Book{})
	sampleBookData, err := GetSampleBookData()
	if err != nil {
		return nil, fmt.Errorf("cannot init Book Repository %v", err)
	}
	for _, book := range sampleBookData {
		bookRepo.db.Where(Book{Id: book.Id}).Attrs(Book{Id: book.Id, Name: book.Name, PageNumber: book.PageNumber,
			Stock: book.Stock, Price: book.Price, StockCode: book.StockCode, ISBN: book.ISBN,
			AuthorId: book.AuthorId}).FirstOrCreate(&book)
	}
	return &bookRepo, nil
}

// GetSampleBookData GetSampleAuthorData reading book json mapping struct and return book list.
func GetSampleBookData() ([]Book, error) {
	var initialBooks []Book
	contents, err := ioutil.ReadFile("./data/book.json")

	if err != nil {
		return nil, fmt.Errorf("cannot read 'book.json' %v", err)
	}
	if err := json.Unmarshal(contents, &initialBooks); err != nil {
		return nil, fmt.Errorf("cannot unmarshall 'book.json' %v", err)
	}
	return initialBooks, nil
}
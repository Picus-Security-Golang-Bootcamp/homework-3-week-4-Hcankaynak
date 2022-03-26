package main

import (
	"fmt"
	"log"
	postgres "main/common/db"
	"main/domain/author"
	"main/domain/book"
)

func main() {
	// loading db
	db, err := postgres.LoadDB()
	if err != nil {
		log.Fatal(err)
	}

	// creating book repository
	bookRepo, err := book.NewBookRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	// creating author repository
	authorRepo, err := author.NewAuthorRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	// Author queries.
	//fmt.Println(authorRepo.FindAll())
	//fmt.Println(authorRepo.FindById(1).ToString())
	//fmt.Println(authorRepo.FindByName("William Shakespeare").ToString())
	authorRepo.FindNameByLike("y")

	// Book queries
	//fmt.Println(bookRepo.FindAll())
	//fmt.Println(bookRepo.FindById(2))
	//fmt.Println(bookRepo.FindByName("Ulysses").ToString())
	//fmt.Println(bookRepo.FindNameByLike("ss"))
	fmt.Println(bookRepo.FindAuthorOfBookById(3, authorRepo).ToString())
}

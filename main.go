package main

import (
	"log"
	postgres "main/common/db"
	"main/domain/book"
)

func main() {

	db, err := postgres.LoadDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = book.NewBookRepository(db)
	if err != nil {
		log.Fatal(err)
	}

	//cityRepo := city.NewCityRepository(db)
	//cityRepo.Migration()
	//cityRepo.InsertSampleData()
	//fmt.Println(cityRepo.FindByCountryCodeWithStruct("01"))

}

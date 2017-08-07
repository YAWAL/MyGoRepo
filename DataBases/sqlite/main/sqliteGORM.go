package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	//"fmt"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	db.Where("code = L1212")

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	query := db.Where("main_user = ? AND contact_user = ?").Find(&product)
	values := query.v

	// Delete - delete product
	db.Delete(&product)
}

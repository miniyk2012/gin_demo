package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type Business struct {
	ID     uint
	Name   string `gorm:"not null"`
	Tables Tables
}

type Table struct {
	ID         uint
	Ref        string `gorm:"not null"`
	Business   Business
	BusinessID uint
}

type Tables []Table
type Businesses []Business

func main() {
	var err error
	var db *gorm.DB

	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&Business{})
	db.AutoMigrate(&Table{})

	bus := Business{
		Name: "Test",
		Tables: Tables{
			Table{Ref: "A1"},
		},
	}
	db.Create(&bus)
	var businesses Businesses
	db.Preload("Tables").Find(&businesses)
	log.Println(businesses)
}
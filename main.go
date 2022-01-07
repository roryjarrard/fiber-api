package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func main() {
	err = connectDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected successfully")
}

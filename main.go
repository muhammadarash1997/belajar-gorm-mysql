package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type DataPerson struct {
	PersonName string
	CarName    string
}

type Person struct {
	ID   int
	Name string
}

type Car struct {
	ID       int
	PersonID int
	Name     string
}

func main() {
	godotenv.Load(".env")
	
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbPass, dbHost, dbPort, dbName)
	// dsn := "root:0000@tcp(127.0.0.1:3306)/trydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Fail connect")
		panic(err.Error())
	}
	fmt.Println("Success connect")

	db.AutoMigrate(Person{})
	db.AutoMigrate(Car{})

	// result := db.Raw("SELECT people.name AS person_name, cars.name AS car_name FROM people INNER JOIN cars ON people.id=cars.person_id WHERE people.name='agus' LIMIT 1")

	// a := DataPerson{}

	// result.Scan(&a)
	// fmt.Println(a)
}

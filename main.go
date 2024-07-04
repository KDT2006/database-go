package main

import (
	"encoding/json"
	"fmt"
)

const Version = "1.0.0"

func main() {
	fmt.Println("Hello World!")
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println(err)
	}

	employees := []User{
		{"John", "23", "123456789", "Tech Inc", Address{"Chennai", "TamilNadu", "India", "600098"}},
		{"Paul", "21", "123456789", "Google", Address{"New York", "New York", "USA", "10001"}},
		{"Jack", "20", "123456789", "Microsoft", Address{"San Francisco", "California", "USA", "12003"}},
		{"Carl", "28", "123456789", "Apple", Address{"New Jersey", "Newark", "USA", "11002"}},
		{"Ben", "32", "123456789", "Nothing", Address{"Louisiana", "New Orleans", "USA", "149932"}},
		{"Eric", "27", "123456789", "Netflix", Address{"Texas", "Austin", "USA", "103441"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error: ", err)
		}
		allusers = append(allusers, employeeFound)
	}

	fmt.Println(allusers)

	// if err := db.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error: ", err)
	// }

	if err := db.Update("users", "Carl", "Age", "48"); err != nil {
		fmt.Println("Error: ", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	Name    string  `json:"name"`
	Id      int     `json:"id"`
	Address Address `json:"address"`
	Dept    string  `json:"dept"`
	Contact Contact `json:"contact"`
}

type Address struct {
	Street  string `json:"street"`
	Houseno int    `json:"houseno"`
	City    string `json:"city"`
}
type Contact struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
}

func main() {
	var filePath string
	fmt.Println("Enter json file path :")
	fmt.Scanln(&filePath)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if !json.Valid(byteValue) {
		fmt.Println("Json file invalid")
		os.Exit(3)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	var students Students

	json.Unmarshal(byteValue, &students)

	for i := 0; i < len(students.Students); i++ {
		fmt.Println()
		fmt.Println("Name : " + students.Students[i].Name)
		fmt.Println("ID : " + strconv.Itoa(students.Students[i].Id))
		fmt.Println("Department : " + students.Students[i].Dept)
		fmt.Println("Address")
		fmt.Println("Street : " + students.Students[i].Address.Street)
		fmt.Println("House number : " + strconv.Itoa(students.Students[i].Address.Houseno))
		fmt.Println("City : " + students.Students[i].Address.City)
		fmt.Println("Contact Numbers")
		fmt.Println("Primary : " + strconv.Itoa(students.Students[i].Contact.Primary))
		fmt.Println("Secondary : " + strconv.Itoa(students.Students[i].Contact.Secondary))
		fmt.Println()
	}

}

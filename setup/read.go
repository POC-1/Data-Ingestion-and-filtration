package setup

import (
	"POC1/layout" //Package for Database layout
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func GetJsonByteVal() []byte {
	var filePath string
	log.Println("Enter json file path :")
	fmt.Scanln(&filePath)
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	if !json.Valid(byteValue) {
		log.Println("Json file invalid")
		os.Exit(3)
	}
	log.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	return byteValue
}

func Display_json(students []layout.Student) {

	for i := 0; i < len(students); i++ {
		log.Println()
		log.Println("Name : " + students[i].Name)
		log.Println("ID : " + strconv.Itoa(students[i].Id))
		log.Println("Department : " + students[i].Dept)
		log.Println("Address")
		log.Println("Street : " + students[i].Address.Street)
		log.Println("House number : " + strconv.Itoa(students[i].Address.Houseno))
		log.Println("City : " + students[i].Address.City)
		log.Println("Contact Numbers")
		log.Println("Primary : " + strconv.Itoa(students[i].Contact.Primary))
		log.Println("Secondary : " + strconv.Itoa(students[i].Contact.Secondary))
		log.Println()
	}
}

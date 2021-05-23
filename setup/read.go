package setup

import (
	"POC1/layout"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func GetJsonByteVal() []byte {
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
	return byteValue
}

func Display_json(students []layout.Student) {

	for i := 0; i < len(students); i++ {
		fmt.Println()
		fmt.Println("Name : " + students[i].Name)
		fmt.Println("ID : " + strconv.Itoa(students[i].Id))
		fmt.Println("Department : " + students[i].Dept)
		fmt.Println("Address")
		fmt.Println("Street : " + students[i].Address.Street)
		fmt.Println("House number : " + strconv.Itoa(students[i].Address.Houseno))
		fmt.Println("City : " + students[i].Address.City)
		fmt.Println("Contact Numbers")
		fmt.Println("Primary : " + strconv.Itoa(students[i].Contact.Primary))
		fmt.Println("Secondary : " + strconv.Itoa(students[i].Contact.Secondary))
		fmt.Println()
	}
}

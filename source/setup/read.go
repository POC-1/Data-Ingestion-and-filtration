package setup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

package read_insert

import (
    "fmt"
    "os"
    "encoding/json"
	"io/ioutil"
	"github.com/POC1/poc_1/models"
    "strconv"
	)
	
func Getfilejson() []byte {
	fmt.Println("Enter file path: ")
	var input_path_json string

	fmt.Scanln(&input_path_json)
	fmt.Print("path for the file : " + input_path_json + "\n")

	// check if its json or not
	jsonData, err := os.Open(input_path_json)
	if err!= nil {
		fmt.Println(err)
		os.Exit(3)
	}
	
	defer jsonData.Close()
	byteValue, _ := ioutil.ReadAll(jsonData)

    if !json.Valid(byteValue) {
		fmt.Println("Json file invalid")
		os.Exit(3)
	}
	return byteValue
}

func Printstudents_docs(students []models.Student) {
	// Print student data
	for i := 0; i < len(students); i++ {
		fmt.Println("\nStudent Name: " + students[i].Name)
		fmt.Println("Student Id: " + strconv.Itoa(students[i].Id))
		fmt.Println("Student Address street: " + students[i].Address.Street)
		fmt.Println("Student Address house no: " + strconv.Itoa(students[i].Address.Houseno))
		fmt.Println("Student Address city: " + students[i].Address.City)
		fmt.Println("Student Dept: " + students[i].Dept)
		fmt.Println("Student Contact Primary: " + strconv.Itoa(students[i].Contact.Primary))
		fmt.Println("Student Contact Secondary: " + strconv.Itoa(students[i].Contact.Secondary))
		fmt.Println()
	}
}

func Getdata_array(students []models.Student) []string{

	// Declare empty array for the document strings
    var docs []string

    for i := 0; i < len(students); i++ {
        // Marshal the struct to JSON and check for errors
        b, err := json.Marshal(students[i])
        if err != nil {
            fmt.Println("json.Marshal ERROR:", err)
            // string(err.Error())
        }
        
        docs = append(docs, string(b))
	}
	 return docs

}
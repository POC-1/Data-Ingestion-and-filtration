package read_insert

import (
	// "fmt"
	"log"
    "os"
    "encoding/json"
	"io/ioutil"
	"github.com/POC1/poc_1/models"
    "strconv"
	)
	
// convert json file data to byte[]
func Getfilejson(input_path_json string) []byte {

	// check if its json or not
	jsonData, err := os.Open(input_path_json)
	if err!= nil {
		log.Println(err)
		return []byte("")
	}
	
	defer jsonData.Close()
	
	byteValue, _ := ioutil.ReadAll(jsonData)

    if !json.Valid(byteValue) {
		log.Println("Json file invalid")
		return []byte("")
	}
	return byteValue
}

// Print data converted to struct from json file 
func Printstudents_docs(students []models.Student) {
	// Print student data
	for i := 0; i < len(students); i++ {
		log.Println("\nStudent Name: " + students[i].Name)
		log.Println("Student Id: " + strconv.Itoa(students[i].Id))
		log.Println("Student Address street: " + students[i].Address.Street)
		log.Println("Student Address house no: " + strconv.Itoa(students[i].Address.Houseno))
		log.Println("Student Address city: " + students[i].Address.City)
		log.Println("Student Dept: " + students[i].Dept)
		log.Println("Student Contact Primary: " + strconv.Itoa(students[i].Contact.Primary))
		log.Println("Student Contact Secondary: " + strconv.Itoa(students[i].Contact.Secondary))
		log.Println()
	}
}

// Makes String array of all the struct records to make it able to insert in elastic
func Getdata_array(students []models.Student) []string{

	// Declare empty array for the document strings
    var docs []string

    for i := 0; i < len(students); i++ {
        // Marshal the struct to JSON and check for errors
        b, err := json.Marshal(students[i])
        if err != nil {
            log.Fatalln("json.Marshal ERROR:", err)
            // string(err.Error())
        }
        
        docs = append(docs, string(b))
	}
	 return docs

}
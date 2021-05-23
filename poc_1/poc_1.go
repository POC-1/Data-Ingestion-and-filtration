package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/POC1/poc_1/models"
	"github.com/POC1/poc_1/query_helper"
	"github.com/POC1/poc_1/read_insert"
)

func main() {

	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Get file path from user and check if it is json
	byteValue := read_insert.Getfilejson()

	var students []models.Student
	json.Unmarshal(byteValue, &students)

	// Print the data from json
	read_insert.Printstudents_docs(students)

	// Get a string array from JSON data
	docs := read_insert.Getdata_array(students)

	// Insert data in elasticsearch
	read_insert.Insert_data(docs)

	var i int = 1
	for i > 0 { // Get query choice
		fmt.Println("Enter Your Choice: \n1. Filter students belonging to a city,'Pune' \n2. Filter students with dept as 'Computer Science' \n3. Filter students with dept as 'Computer Application' \n4. Filter students with dept containing 'Computer' \n5. END ")
		var input_query_type int

		fmt.Scanln(&input_query_type)
		fmt.Print("Query selected : " + strconv.Itoa(input_query_type) + "\n")

		var query = ``
		var score = 1
		read := strings.NewReader("")
		switch input_query_type {
		case 1:
			query = `"match" : {"dept": "Computer Science"}`
			read = query_helper.ConstructNestedQuery()
		case 2:
			query = `"match" : {"dept": "Computer Science"}`
			score = 1
			read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
		case 3:
			query = `"match" : {"dept": "Computer Application"}`
			score = 1
			read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
		case 4:
			query = `"match" : {"dept": "Computer"}`
			score = 0
			read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
		case 5:
			fmt.Print("ENDING..... ")
			os.Exit(3)
		default:
			fmt.Print("Wrong Input!! ")
			os.Exit(3)
		}

		// fmt.Println("read:", read)
		fmt.Println("read TYPE:", reflect.TypeOf(read))
		// fmt.Println("JSON encoding:", json.NewEncoder(&buf).Encode(read))

		// Call Query Function
		query_helper.Makequery(read)
	}

}

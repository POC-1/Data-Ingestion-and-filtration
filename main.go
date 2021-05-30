package main

import (
	es_services "POC1/elasticsearch_services" //Package for Elastic Services
	"POC1/layout"                             //Package for Database Structure
	"POC1/setup"                              //Package for reading & printing Json file
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var filePath string
	log.Println("Enter json file path :")
	fmt.Scanln(&filePath)
	byteValue := setup.GetJsonByteVal(filePath) //Read and Verify Json File
	if bytes.Compare(byteValue, []byte("")) == 0 {
		os.Exit(3)
	} //Check if it is a valid byte format

	var students []layout.Student
	json.Unmarshal(byteValue, &students)

	setup.Display_json(students) // Print Json File Content

	es_client := es_services.SetClient()                  // Instantiate a new Elasticsearch client object instance
	ctx := es_services.InsertRecords(students, es_client) // Insert Records

	var query string
	var ch int
	var i int = 1
	for i > 0 {

		log.Println("Enter your choice :\n1.Find students belonging to the city 'Pune'")
		log.Println("2.Find students belonging to the department 'Computer Science'")
		log.Println("3.Find students belonging to the department 'Computer Application'")
		log.Println("4.Find students with department containing 'Computer' ")
		log.Println("5.Display all students ")
		log.Println("6.Exit")

		switch fmt.Scanln(&ch); ch {
		case 1:
			query = `{
				"query": {
				  "nested": {
					"path": "address",
					"query": {
					  "bool": {
						"must": [
						  { "match": { "address.city": "Pune" } }
						]
					  }
					},
					"score_mode": "avg"
				  }
				}
			  }`

		case 2:
			query = `{
				"min_score": 1,
				"query" : {
					"match" : { "dept" : "Computer Science" }
				}
			}`
		case 3:
			query = `{
				"min_score": 1,
				"query" : {
					"match" : { "dept" : "Computer Application" }
				}
			}`
		case 4:
			query = `{
				"query" : {
					"match" : { "dept" : "Computer" }
				}
			}`
		case 5:
			query = `{
				"query": { 
				  "match_all": {}
				}
			  }`

		case 6:
			log.Println("Exiting...")
			os.Exit(3)

		default:
			log.Println("Invalid choice")
		}

		// Pass the query string to the function and have it return a Reader object
		read := es_services.ConstructQuery(query)
		if ch == 0 {
			log.Println("Using default match_all query")
		}
		es_services.CallQuery(es_client, read, ctx) //Print results

	} //End of loop

}

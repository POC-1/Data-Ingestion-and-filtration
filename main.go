package main

import (
	es_services "POC1/elasticsearch_services"
	"POC1/layout"
	"POC1/setup"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	byteValue := setup.GetJsonByteVal() //Read and Verify Json File

	var students []layout.Student
	json.Unmarshal(byteValue, &students)

	setup.Display_json(students) // Print Json File Content

	es_client := es_services.SetClient()                  // Instantiate a new Elasticsearch client object instance
	ctx := es_services.InsertRecords(students, es_client) // Insert Records

	var query string
	var ch int
	var i int = 1
	for i > 0 {

		fmt.Println("Enter your choice :\n1.Find students belonging to the city 'Pune'")
		fmt.Println("2.Find students belonging to the department 'Computer Science'")
		fmt.Println("3.Find students belonging to the department 'Computer Application'")
		fmt.Println("4.Find students with department containing 'Computer' ")
		fmt.Println("5.Display all students ")
		fmt.Println("6.Exit")

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
			fmt.Println("Exiting...")
			os.Exit(3)

		default:
			fmt.Println("Invalid choice")
		}

		// Pass the query string to the function and have it return a Reader object
		read := es_services.ConstructQuery(query)
		es_services.CallQuery(es_client, read, ctx) //Print results

	} //End of while loop

}

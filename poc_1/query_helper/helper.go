package query_helper

import (
    "fmt"
	"encoding/json"
	"strings"
	"strconv"
	"log"
	)
	
// Asks for size of documents to get as result
func Get_resultsize() int{
	log.Println("Enter no of allowed documents to get: ")
	var size_expected int
	fmt.Scanln(&size_expected)
	log.Print("Query selected : " + strconv.Itoa(size_expected) + "\n")
	return size_expected
}

// Converts query string to strings.Reader
func makeReader(query string) *strings.Reader {
	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

// Constructs query for case 2, 3, 4
func ConstructQuery(q string, size int, score int) *strings.Reader {

	// Build a query string from string passed to function
	var query = `{"min_score":` + strconv.Itoa(score) + `, "query": {`
	
	// Concatenate query string with string passed to method call
	query = query + q
	
	// Use the strconv.Itoa() method to convert int to string
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`
	log.Println("\nquery:", query)
	
	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool
	
	// Default query is "{}" if JSON is invalid
	if isValid == false {
	log.Println("constructQuery() ERROR: query string not valid:", query)
	log.Println("Using default match_all query")
	query = "{}"
	} else {
		log.Println("constructQuery() valid JSON:", isValid)
	}
		
	return makeReader(query)
}

// Constructs query for case 1
func ConstructNestedQuery() *strings.Reader {

		// Build a query string from string passed to function
		var query = `{
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

		  log.Println("\nquery:", query)
		
		// Check for JSON errors
		isValid := json.Valid([]byte(query)) // returns bool
		
		// Default query is "{}" if JSON is invalid
		if isValid == false {
		log.Println("constructQuery() ERROR: query string not valid:", query)
		log.Println("Using default match_all query")
		query = "{}"
		} else {
			log.Println("constructQuery() valid JSON:", isValid)
		}
			
		return makeReader(query)
}

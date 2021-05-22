package query_helper

import (
    "fmt"
	"encoding/json"
	"strings"
	"strconv"
	)
	
func Get_resultsize() int{
	fmt.Println("Enter no of allowed documents to get: ")
	var size_expected int
	fmt.Scanln(&size_expected)
	fmt.Print("Query selected : " + strconv.Itoa(size_expected) + "\n")
	return size_expected
}

func makeReader(query string) *strings.Reader {
	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

func ConstructQuery(q string, size int, score int) *strings.Reader {

	// Build a query string from string passed to function
	var query = `{"min_score":` + strconv.Itoa(score) + `, "query": {`
	
	// Concatenate query string with string passed to method call
	query = query + q
	
	// Use the strconv.Itoa() method to convert int to string
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`
	fmt.Println("\nquery:", query)
	
	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool
	
	// Default query is "{}" if JSON is invalid
	if isValid == false {
	fmt.Println("constructQuery() ERROR: query string not valid:", query)
	fmt.Println("Using default match_all query")
	query = "{}"
	} else {
	fmt.Println("constructQuery() valid JSON:", isValid)
	}
		
	return makeReader(query)
}

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

		fmt.Println("\nquery:", query)
		
		// Check for JSON errors
		isValid := json.Valid([]byte(query)) // returns bool
		
		// Default query is "{}" if JSON is invalid
		if isValid == false {
		fmt.Println("constructQuery() ERROR: query string not valid:", query)
		fmt.Println("Using default match_all query")
		query = "{}"
		} else {
		fmt.Println("constructQuery() valid JSON:", isValid)
		}
			
		return makeReader(query)
}

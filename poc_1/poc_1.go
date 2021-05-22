package main 

import (
    "fmt"
	"os"
	"bytes"
	"log"
	"encoding/json"
    // "path/filepath"
    "github.com/POC1/poc_1/util"
    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
    // "github.com/elastic/go-elasticsearch/v7/esapi"
    "context"
    "reflect"
	"strings"
	"strconv"
    )


type Student struct {
    Name        string  `json:"name"`
    Id          int  `json:"id"`
    Address     Address `json:"address"`
    Dept        string  `json:"dept"`
    Contact     Contact `json:"contact"`
}

type Address struct{
    Street        string  `json:"street"`
    Houseno       int  `json:"houseno"`
    City          string  `json:"city"`
}

type Contact struct{
    Primary        int  `json:"primary"`
    Secondary      int  `json:"secondary"`
}

func check_file_path(input_path_json string) bool{
    if _, err := os.Stat(input_path_json); os.IsNotExist(err) {
        println("\nNo Such file on given path")
        return false
    }
    println("\n\nFile Exists! ")
    return true
}

func constructQuery(q string, size int, score int) *strings.Reader {

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
		
	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
	}

	func constructNestedQuery() *strings.Reader {

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
			
		// Build a new string from JSON query
		var b strings.Builder
		b.WriteString(query)
	
		// Instantiate a *strings.Reader object from string
		read := strings.NewReader(b.String())
	
		// Return a *strings.Reader object
		return read
	}

func main() {

	// Allow for custom formatting of log output
	log.SetFlags(0)

    // Load Config variables
    config, err := util.LoadConfig(".")
    if err != nil {
        // log.Fatal("Cannot load config: ", err)
        fmt.Println("Cannot load config: ", err)
    }

    // Create a context object for the API calls
    ctx := context.Background()

    // Declare an Elasticsearch configuration
    cfg := elasticsearch7.Config{
        Addresses: []string{
            config.ELASTICSEARCH_URL,
        },
        // Username: config.USERNAME,
        // Password: config.PASSWORD,
    }

    // Instantiate a new Elasticsearch client object instance
    client, err := elasticsearch7.NewClient(cfg)

    if err != nil {
        fmt.Println("Elasticsearch connection error:", err)
    }

    // Have the client instance return a response
    res, err := client.Info()

    // Deserialize the response into a map.
    if err != nil {
        log.Fatalf("client.Info() ERROR:", err)
    } else {
        log.Printf("client response:", res)
    }


	 // Get file path from user 
	 fmt.Println("Enter Your Choice: \n1. Filter students belonging to a a city,'Pune' \n2. Filter students with dept as 'Computer Science' \n3. Filter students with dept as 'Computer Application' \n4. Filter students with dept containing 'Computer' ")
	 var input_query_type int
 
	 fmt.Scanln(&input_query_type)
	 fmt.Print("Query selected : " + strconv.Itoa(input_query_type) + "\n")


	var query = ``
	var score = 1
	read := strings.NewReader("")
	switch input_query_type {
    case 1:
        query = `"match" : {"dept": "Computer Science"}`
		read = constructNestedQuery()
	case 2:
		fmt.Println("Enter no of allowed documents to get: ")
		var size_expected int
		fmt.Scanln(&size_expected)
		fmt.Print("Query selected : " + strconv.Itoa(size_expected) + "\n")

        query = `"match" : {"dept": "Computer Science"}`
		score = 1
		read = constructQuery(query, size_expected, score)
	case 3:
		fmt.Println("Enter no of allowed documents to get: ")
		var size_expected int
		fmt.Scanln(&size_expected)
		fmt.Print("Query selected : " + strconv.Itoa(size_expected) + "\n")

		query = `"match" : {"dept": "Computer Application"}`
		score = 1
		read = constructQuery(query, size_expected, score)
	case 4:
		fmt.Println("Enter no of allowed documents to get: ")
		var size_expected int
		fmt.Scanln(&size_expected)
		fmt.Print("Query selected : " + strconv.Itoa(size_expected) + "\n")

		query = `"match" : {"dept": "Computer"}`
		score = 0
		read = constructQuery(query, size_expected, score)
	default:
		fmt.Print("Wrong Input, Printing default i.e 2")
        os.Exit(3)
	}
	
	// Instantiate a mapping interface for API response
	var mapResp map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer
	
	// size = size_expected


	// fmt.Println("read:", read)
	fmt.Println("read TYPE:", reflect.TypeOf(read))
	// fmt.Println("JSON encoding:", json.NewEncoder(&buf).Encode(read))

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
	log.Fatalf("Error encoding query: %s", err)

	// Query is a valid JSON object
	} else {
		fmt.Println("json.NewEncoder encoded query:", read)
		fmt.Println()

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
		client.Search.WithContext(ctx),
		client.Search.WithIndex("poc_one_t"),
		client.Search.WithBody(read),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
		log.Fatalf("Elasticsearch Search() API ERROR:", err)
		
		// If no errors are returned, parse esapi.Response object
		} else {
		fmt.Println("res TYPE:", reflect.TypeOf(res))
		
		// Close the result body when the function call is complete
		defer res.Body.Close()

		// Decode the JSON response and using a pointer
		if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {
		// fmt.Println(`&mapResp:`, &mapResp, "\n")
		// fmt.Println(`mapResp["hits"]:`, mapResp["hits"])

		// Iterate the document "hits" returned by API call
		for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

			// Parse the attributes/fields of the document
			doc := hit.(map[string]interface{})
			
			// The "_source" data is another map interface nested inside of doc
			source := doc["_source"]
			fmt.Println("doc _source:", reflect.TypeOf(source))
			
			// Get the document's _id and print it out along with _source data
			docID := doc["_id"]
			fmt.Println("docID:", docID)
			fmt.Println("_source:", source)
			fmt.Println()
		} // end of response iteration
			
		}
		}
	}
		
}
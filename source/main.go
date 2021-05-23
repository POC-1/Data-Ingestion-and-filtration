package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	elastic_search "github.com/elastic/go-elasticsearch/v7"
	elastic "github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string  `json:"name"`
	Id      int     `json:"id"`
	Address Address `json:"address"`
	Dept    string  `json:"dept"`
	Contact Contact `json:"contact"`
}

type Address struct {
	Street  string `json:"street"`
	Houseno int    `json:"houseno"`
	City    string `json:"city"`
}
type Contact struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
}

func main() {

	/*
		var students []Student

		json.Unmarshal(byteValue, &students)

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



		//SetClient()
		// Instantiate a new Elasticsearch client object instance



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

		for i, bod := range docs {

			fmt.Println("\nDOC _id:", i+1)
			fmt.Println(bod)

			// Instantiate a request object
			req := esapi.IndexRequest{
				Index:      "students",
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(bod),
				Refresh:    "true",
			}
			fmt.Println(reflect.TypeOf(req))

			// Return an API response object from request
			res, err := req.Do(ctx, client)
			if err != nil {
				log.Fatalf("IndexRequest ERROR: %s", err)
			}
			defer res.Body.Close()
			fmt.Printf("res val %s", res)
			if res.IsError() {
				log.Printf("%s ERROR indexing document ID=%d", res.Status(), i+1)
			} else {

				// Deserialize the response into a map.
				var resMap map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					log.Printf("\nIndexRequest() RESPONSE:")
					// Print the response status and indexed document version.
					fmt.Println("Status:", res.Status())
					fmt.Println("Result:", resMap["result"])
					fmt.Println("Version:", int(resMap["_version"].(float64)))
					fmt.Println("resMap:", resMap)
					fmt.Println()

				}
			}

		} */
	//var query = `{"query": {"match" : {"dept": "Computer"}},"size": 10}`
	cfg := elastic_search.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// Username: config.USERNAME,
		// Password: config.PASSWORD,
	}
	//Create a context object for the API calls
	ctx := context.Background()

	client, err := elastic_search.NewClient(cfg)

	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}
	// Have the client instance return a response
	res, err := client.Info()

	// Deserialize the response into a map.
	if err != nil {
		fmt.Println("client.Info() ERROR:", err)
		os.Exit(3)
	} else {
		fmt.Println("client response:", res)
	}
	var query string
	var ch int
	fmt.Println("Enter your choice :\n1.Find students belonging to the city 'Pune'")
	fmt.Println("2.Find students belonging to the department 'Computer Science'")
	fmt.Println("3.Find students belonging to the department 'Computer Application'")
	fmt.Println("4.Find students with department containing 'Computer' ")
	fmt.Println("5.Display all students ")

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
	default:
		fmt.Println("Invalid choice")
	}
	// Pass the query string to the function and have it return a Reader object
	read := constructQuery(query, 2)
	fmt.Println("read:", read)
	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)
	} else {
		fmt.Println("json.NewEncoder encoded query:", read, "\n")

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex("students"),
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

				for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

					// Parse the attributes/fields of the document
					doc := hit.(map[string]interface{})

					// The "_source" data is another map interface nested inside of doc
					source := doc["_source"]
					//fmt.Println("doc _source:", reflect.TypeOf(source))

					// Get the document's _id and print it out along with _source data
					docID := doc["_id"]
					fmt.Println("docID:", docID)
					fmt.Println("_source:", source)
					val := source.(map[string]interface{})
					fmt.Println("ID :", val["id"])
					fmt.Println("Name :", val["name"])
					fmt.Println("Department :", val["dept"])
					fmt.Println("Address:", val["address"])
					fmt.Println("Contact :", val["contact"])

					fmt.Println()
					fmt.Println()
				} // end of response iteration
			}
		}
	}

}

func SetClient() {
	es, err := elastic_search.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func constructQuery(query string, size int) *strings.Reader {
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

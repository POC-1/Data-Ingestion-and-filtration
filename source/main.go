package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	elastic_search "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	elastic "github.com/olivere/elastic/v7"
)

type Students struct {
	Students []Student `json:"students"`
}

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

	var students Students

	json.Unmarshal(byteValue, &students)

	for i := 0; i < len(students.Students); i++ {
		fmt.Println()
		fmt.Println("Name : " + students.Students[i].Name)
		fmt.Println("ID : " + strconv.Itoa(students.Students[i].Id))
		fmt.Println("Department : " + students.Students[i].Dept)
		fmt.Println("Address")
		fmt.Println("Street : " + students.Students[i].Address.Street)
		fmt.Println("House number : " + strconv.Itoa(students.Students[i].Address.Houseno))
		fmt.Println("City : " + students.Students[i].Address.City)
		fmt.Println("Contact Numbers")
		fmt.Println("Primary : " + strconv.Itoa(students.Students[i].Contact.Primary))
		fmt.Println("Secondary : " + strconv.Itoa(students.Students[i].Contact.Secondary))
		fmt.Println()
	}

	/*

		js := string(byteValue)
		ind, err := esclient.Index().
			Index("students").
			BodyJson(js).
			Do(ctx)

		if err != nil {
			panic(err)
		}

		fmt.Println("[Elastic][InsertProduct]Insertion Successful" + ind.Index)*/
	//SetClient()
	// Instantiate a new Elasticsearch client object instance

	cfg := elastic_search.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// Username: config.USERNAME,
		// Password: config.PASSWORD,
	}

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
	// Create a context object for the API calls
	ctx := context.Background()

	// Declare empty array for the document strings
	var docs []string

	for i := 0; i < len(students.Students); i++ {
		// Marshal the struct to JSON and check for errors
		b, err := json.Marshal(students.Students[i])
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

	}
	//var query = `{"query": {"match" : {"dept": "Computer"}},"size": 10}`
	var query = `"bool": {
		"filter": {
		"match": {
		"dept" : "Computer"
				}}}`
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
				fmt.Println(`&mapResp:`, &mapResp)
				fmt.Println(`mapResp["hits"]:`, mapResp["hits"])
				//fmt.Println(`Result:`, mapResp["result"])
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

func constructQuery(q string, size int) *strings.Reader {
	// Build a query string from string passed to function
	var query = `{"query": {`

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

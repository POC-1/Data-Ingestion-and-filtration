package elasticsearch_services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"reflect"
	"strings"

	elastic_search "github.com/elastic/go-elasticsearch/v7"
)

func ConstructQuery(query string) *strings.Reader {
	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool

	// Default query is "{}" if JSON is invalid
	if isValid == false {
		log.Println("constructQuery() ERROR: query string not valid:", query)
		query = "{}"
	} else {
		log.Println("constructQuery() valid JSON:", isValid)
	}
	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

func CallQuery(client *elastic_search.Client, read *strings.Reader, ctx context.Context) {
	// Instantiate a map interface object for storing returned documents
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Println("json.NewEncoder() ERROR:", err)
		os.Exit(3)
	} else {
		log.Println("json.NewEncoder encoded query:", read)

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
			log.Println("Elasticsearch Search() API ERROR:", err)
			os.Exit(3)

			// If no errors are returned, parse esapi.Response object
		} else {
			log.Println("res TYPE:", reflect.TypeOf(res))

			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err == nil {

				for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

					// Parse the attributes/fields of the document
					doc := hit.(map[string]interface{})

					// The "_source" data is another map interface nested inside of doc
					source := doc["_source"]
					//log.Println("doc _source:", reflect.TypeOf(source))

					// Get the document's _id and print it out along with _source data
					docID := doc["_id"]
					log.Println("docID:", docID)
					log.Println("_source:", source)
					val := source.(map[string]interface{})
					log.Println("ID :", val["id"])
					log.Println("Name :", val["name"])
					log.Println("Department :", val["dept"])
					log.Println("Address:", val["address"])
					log.Println("Contact :", val["contact"])

					log.Println()
					log.Println()
				} // end of response iteration
			}
		}
	}
}

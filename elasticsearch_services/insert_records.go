package elasticsearch_services

import (
	"POC1/layout"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	elastic_search "github.com/elastic/go-elasticsearch/v7"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func InsertRecords(students []layout.Student, client *elastic_search.Client) context.Context {
	// Declare empty array for the document strings
	var docs []string

	for i := 0; i < len(students); i++ {
		// Marshal the struct to JSON and check for errors
		b, err := json.Marshal(students[i])
		if err != nil {
			fmt.Println("json.Marshal ERROR:", err)
		}

		docs = append(docs, string(b))
	}
	ctx := context.Background()

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
	return ctx
}

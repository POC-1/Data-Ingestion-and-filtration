package query_helper

import (
	"bytes"
	"log"
	"encoding/json"
    "github.com/POC1/poc_1/util"
    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
    // "github.com/elastic/go-elasticsearch/v7/esapi"
    "context"
    "reflect"
	"strings"
	)

// Logic for querying in elasticsearch
func Makequery(read *strings.Reader) {

	// Load Config variables
	config, err := util.LoadConfig(".")
	if err != nil {
		// log.Fatal("Cannot load config: ", err)
		log.Println("Cannot load config: ", err)
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
        log.Println("Elasticsearch connection error:", err)
    }

    // Have the client instance return a response
    res, err := client.Info()

    // Deserialize the response into a map.
    if err != nil {
        log.Fatalf("client.Info() ERROR:", err)
    } else {
        log.Printf("client response:", res)
    }

	// Instantiate a mapping interface for API response
	var mapResp map[string]interface{}

	// Build the request body.
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	} else {
		// Query is a valid JSON object 

		log.Println("json.NewEncoder encoded query:", read)
		log.Println()

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex(config.INDEX_NAME),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
			client.Search.WithPretty(),
			)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
		log.Fatalf("Elasticsearch Search() API ERROR:", err)
		
		// If no errors are returned, parse esapi.Response object
		} else {
			log.Println("res TYPE:", reflect.TypeOf(res))
			
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
					log.Println("doc _source:", reflect.TypeOf(source))
					
					// Get the document's _id and print it out along with _source data
					docID := doc["_id"]
					log.Println("docID:", docID)
					log.Println("_source:", source)
					log.Println()
				} // end of response iteration
				
			}
		}
	}
}

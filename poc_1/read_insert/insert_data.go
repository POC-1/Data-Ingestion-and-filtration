package read_insert 

import (
    "fmt"
    "encoding/json"
    "strconv"
    "log"
    "github.com/POC1/poc_1/util"
    "github.com/POC1/poc_1/go_elastic"
    "github.com/elastic/go-elasticsearch/v7/esapi"
    "context"
    "reflect"
    "strings"
	)
	

// Insert data logic for elasticsearch
func Insert_data(docs []string){


    // Create a context object for the API calls
    ctx := context.Background()

    // Create a mapping for the Elasticsearch documents
    var (
        docMap map[string]interface{}
    )
    log.Println("docMap:", docMap)
    log.Println("docMap TYPE:", reflect.TypeOf(docMap))
	
	// Create coonection and get client
	client := go_elastic.GetClient()
	
    // Load Config variables
	config, err := util.LoadConfig(".")
	if err != nil {
		// log.Fatal("Cannot load config: ", err)
		log.Println("Cannot load config: ", err)
    }
    
	// Iterate on docs and insert records in elastic
	for i, bod := range docs {

		log.Println("\nDOC _id:", i+1)
		log.Println(bod)

		// Instantiate a request object 
		req := esapi.IndexRequest{
			Index:      config.INDEX_NAME,
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(bod),
			Refresh:    "true",
		}
		log.Println(reflect.TypeOf(req))

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
				log.Println("Status:", res.Status())
				log.Println("Result:", resMap["result"])
				log.Println("Version:", int(resMap["_version"].(float64)))
				log.Println("resMap:", resMap)
				log.Println()
			
			}
		}
	} 
}
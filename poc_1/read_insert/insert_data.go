package read_insert 

import (
    "fmt"
    "encoding/json"
    "strconv"
    "log"
    "github.com/POC1/poc_1/util"
    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
    "github.com/elastic/go-elasticsearch/v7/esapi"
    "context"
    "reflect"
    "strings"
	)
	

func Insert_data(docs []string){
	// Load Config variables
    config, err := util.LoadConfig(".")
    if err != nil {
        // log.Fatal("Cannot load config: ", err)
        fmt.Println("Cannot load config: ", err)
    }

    // Create a context object for the API calls
    ctx := context.Background()

    // Create a mapping for the Elasticsearch documents
    var (
        docMap map[string]interface{}
    )
    fmt.Println("docMap:", docMap)
    fmt.Println("docMap TYPE:", reflect.TypeOf(docMap))

    // Declare an Elasticsearch configuration
    cfg := elasticsearch7.Config{
        Addresses: []string{
            config.ELASTICSEARCH_URL,
        },
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
	
	// Iterate on docs and insert records in elastic
	for i, bod := range docs {

		fmt.Println("\nDOC _id:", i+1)
		fmt.Println(bod)

		// Instantiate a request object 
		req := esapi.IndexRequest{
			Index:      config.INDEX_NAME,
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
}
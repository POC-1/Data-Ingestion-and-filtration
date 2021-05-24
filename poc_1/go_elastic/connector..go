package go_elastic

import (
	"github.com/POC1/poc_1/util"
	"log"
	elastic "github.com/elastic/go-elasticsearch/v7"
)

func GetClient() *elastic.Client {

	// Load Config variables
	config, err := util.LoadConfig(".")
	if err != nil {
		// log.Fatal("Cannot load config: ", err)
		log.Println("Cannot load config: ", err)
	}

	// Declare an Elasticsearch configuration
	cfg := elastic.Config{
		Addresses: []string{
			config.ELASTICSEARCH_URL,
		},
	}
	
	
    // Instantiate a new Elasticsearch client object instance
	client, err := elastic.NewClient(cfg)

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

	return client
}
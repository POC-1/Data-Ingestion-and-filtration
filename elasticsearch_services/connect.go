package elasticsearch_services

import (
	"POC1/util"
	"fmt"
	"os"

	elastic_search "github.com/elastic/go-elasticsearch/v7"
)

func SetClient() *elastic_search.Client {

	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load config: ", err)
		os.Exit(3)
	}

	cfg := elastic_search.Config{
		Addresses: []string{
			config.ELASTICSEARCH_URL,
		},
		// Username: config.USERNAME,
		// Password: config.PASSWORD,
	}
	//Create a context object for the API calls
	//ctx := context.Background()

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
	return client
}

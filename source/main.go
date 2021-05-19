package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	elastic_search "github.com/elastic/go-elasticsearch/v7"
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

// Mapping and index name. Elasticsearch index doctypes now deprecated
const (
	index    = "students"
	mappings = `
	{
	"settings":{
	"number_of_shards":2,
	"number_of_replicas":1
	},
	"mappings":{
	"properties":{
	"field str":{
	"type":"text"
	},
	"field int":{
	"type":"integer"
	},
	"field bool":{
	"type":"boolean"
	}
	}
	}
	}
	`
)

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

	//SetClient()
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	js := string(byteValue)
	ind, err := esclient.Index().
		Index("students").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful" + ind.Index)

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

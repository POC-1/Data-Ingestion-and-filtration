package main 

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "log"
    // "path/filepath"
    "github.com/POC1/poc_1/util"

    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
    "github.com/elastic/go-elasticsearch/v7/esapi"
    "context"
    "reflect"
    "strings"
    )

// type Students struct {
//     Student []Student 
// }



type Student struct {
    Name        string  `json:"name"`
    Id          int  `json:"id"`
    Address     Address `json:"address"`
    Dept        string  `json:"dept"`
    Contact     Contact `json:"contact"`
}

type Address struct{
    Street        string  `json:"street"`
    Houseno       int  `json:"houseno"`
    City          string  `json:"city"`
}

type Contact struct{
    Primary        int  `json:"primary"`
    Secondary      int  `json:"secondary"`
}

func check_file_path(input_path_json string) bool{
    if _, err := os.Stat(input_path_json); os.IsNotExist(err) {
        println("\nNo Such file on given path")
        return false
    }
    println("\n\nFile Exists! ")
    return true
}

func main() {

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
        // Username: config.USERNAME,
        // Password: config.PASSWORD,
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

    
    // Get file path from user 
    fmt.Println("Enter file path: ")
    var input_path_json string

    fmt.Scanln(&input_path_json)
    fmt.Print("path for the file : " + input_path_json + "\n")

    // check if its json or not
    jsonData, err := os.Open(input_path_json)
    if err!= nil {
        fmt.Println(err)
        os.Exit(3)
    }
    
    defer jsonData.Close()

    byteValue, _ := ioutil.ReadAll(jsonData)

    if !json.Valid(byteValue) {
		fmt.Println("Json file invalid")
		os.Exit(3)
    }



    var students []Student

    json.Unmarshal(byteValue, &students)

    // Print student data
    for i := 0; i < len(students); i++ {
        fmt.Println("\nStudent Name: " + students[i].Name)
        fmt.Println("Student Id: " + strconv.Itoa(students[i].Id))
        fmt.Println("Student Address street: " + students[i].Address.Street)
        fmt.Println("Student Address house no: " + strconv.Itoa(students[i].Address.Houseno))
        fmt.Println("Student Address city: " + students[i].Address.City)
        fmt.Println("Student Dept: " + students[i].Dept)
        fmt.Println("Student Contact Primary: " + strconv.Itoa(students[i].Contact.Primary))
        fmt.Println("Student Contact Secondary: " + strconv.Itoa(students[i].Contact.Secondary))
        fmt.Println()
    }

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

    // Iterate on docs and insert records in elastic
    for i, bod := range docs {

        fmt.Println("\nDOC _id:", i+1)
        fmt.Println(bod)

        // Instantiate a request object
        req := esapi.IndexRequest{
            Index:      "poc_one_t",
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
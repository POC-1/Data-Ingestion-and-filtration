package main 

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
    "strconv"
    // "path/filepath"
    "github.com/POC1/poc_1/util"
    elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
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

    es, _ := elasticsearch7.NewDefaultClient()
    fmt.Println(elasticsearch7.Version)
    fmt.Println(es.Info())

    es, err := elasticsearch7.NewDefaultClient()
    if err != nil {
    fmt.Printf("Error creating the client: %s", err)
    }

    res, err := es.Info()
    if err != nil {
    fmt.Printf("Error getting response: %s", err)
    }

    defer res.Body.Close()
    fmt.Println(res)

    // Reading variable from config file
    config, err := util.LoadConfig(".")
    if err != nil {
        // log.Fatal("Cannot load config: ", err)
        fmt.Println("Cannot load config: ", err)
    }
    fmt.Println(config.ELASTICSEARCH_URL)

    fmt.Println("Enter file path: ")
    var input_path_json string

    fmt.Scanln(&input_path_json)
    fmt.Print("path for the file : " + input_path_json + "\n")

    // get extension for the file
    // if check_file_path(input_path_json) {
    //     println("Checking if its JSON..")
    // }

    jsonData, err := os.Open(input_path_json)
    if err!= nil {
        fmt.Println(err)
        os.Exit(3)
    }

    // extension := filepath.Ext(input_path_json)
    // if extension != ".json" {
    //     fmt.Println("\nCan't proceed, Extension of file is different! ", extension)
    // }
    
   
    
    defer jsonData.Close()

    byteValue, _ := ioutil.ReadAll(jsonData)

    if !json.Valid(byteValue) {
		fmt.Println("Json file invalid")
		os.Exit(3)
    }

    var students []Student

    json.Unmarshal(byteValue, &students)

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

}

// C:/Users/admin/Desktop/poc1/poc_1/sample.json
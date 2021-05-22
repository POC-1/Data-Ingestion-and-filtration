package main 

import (
    "fmt"
	"os"
	"log"
	"github.com/POC1/poc_1/query_helper"
    "reflect"
	"strings"
	"strconv"
    )


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

func main() {

	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Get query choice
	fmt.Println("Enter Your Choice: \n1. Filter students belonging to a city,'Pune' \n2. Filter students with dept as 'Computer Science' \n3. Filter students with dept as 'Computer Application' \n4. Filter students with dept containing 'Computer' ")
	var input_query_type int

	fmt.Scanln(&input_query_type)
	fmt.Print("Query selected : " + strconv.Itoa(input_query_type) + "\n")

	var query = ``
	var score = 1
	read := strings.NewReader("")
	switch input_query_type {
    case 1:
        query = `"match" : {"dept": "Computer Science"}`
		read = query_helper.ConstructNestedQuery()
	case 2:
        query = `"match" : {"dept": "Computer Science"}`
		score = 1
		read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
	case 3:
		query = `"match" : {"dept": "Computer Application"}`
		score = 1
		read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
	case 4:
		query = `"match" : {"dept": "Computer"}`
		score = 0
		read = query_helper.ConstructQuery(query, query_helper.Get_resultsize(), score)
	default:
		fmt.Print("Wrong Input!! ")
        os.Exit(3)
	}
	
	// fmt.Println("read:", read)
	fmt.Println("read TYPE:", reflect.TypeOf(read))
	// fmt.Println("JSON encoding:", json.NewEncoder(&buf).Encode(read))

	// Call Query Function
	query_helper.Makequery(read)

	
		
}
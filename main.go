package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BryanBD16/smallListApp/list"
)

func main() {
	// Replace with your MySQL DSN (username:password@tcp(host:port)/dbname)
	dsn := "root:1773548_Nype58@tcp(127.0.0.1:3306)/list_app" //Ip address local host
	repo, err := list.NewRepository(dsn)
	if err != nil {
		//The log.Fatal function in Go is part of the standard log package and is used to log an error message and terminate the program immediately.
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	//The Close method ensures that all resources (e.g., database connections) used by the Repository are properly released when you are done with them
	defer repo.Close()
	s := list.NewService(repo)

	//net/http package
	//register a handler function for a specific HTTP route
	http.HandleFunc("/element", s.Get)
	http.HandleFunc("/element/add", s.Add)
	http.HandleFunc("/element/clear", s.Clear)

	fmt.Println("serving port 3000")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}

// example of get method request in terminal
//curl http://localhost:3000/element
//example of add method request in terminal
//curl -X POST -H "Content-Type: application/json" -d '{"Name":"test item","Description":"This is a test item"}' http://localhost:3000/element/add
// example of clear method request in terminal
//curl http://localhost:3000/element/clear
// command to see all database
//mysql -u root -p -e 'SHOW DATABASES'

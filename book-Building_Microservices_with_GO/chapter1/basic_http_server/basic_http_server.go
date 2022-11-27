package main

import (
	"fmt"
	"log"
	"net/http"
)

// 실행방법 : http://localhost:8080/helloworld
func main() {
	port := 8080

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World\n")
	})

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// 원래 책 내용
// func main() {
// 	port := 8080

// 	http.HandleFunc("/helloworld", helloWorldHandler)

// 	log.Printf("Server starting on port %v\n", port)
// 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
// }

// func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World\n")
// }

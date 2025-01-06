package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err:=r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error: %v", err)
		return
	}
	fmt.Fprintf(w, "Form parsed successfully\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	other_info := r.FormValue("other_info")
	fmt.Fprintf(w, "Name = %s\nAddress = %s\nOther info = %s\n", name, address, other_info)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Not Acceptable", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err:=http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
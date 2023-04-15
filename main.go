package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// first this will handle the get request with the html file and then after the data is recieved it will handle the post request with correct values
	switch r.Method {
	case "GET":
		 http.ServeFile(w, r, "./static/form.html")
	case "POST":

		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v\n", err)
			return
		}

		// in case of successful POST request
		fmt.Fprintf(w, "POST request successful")

		// saving the form data in variables
		name := r.FormValue("name")
		address := r.FormValue("address")

		// printing the variables
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// in case of helloHandler being called from a different path
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// in case of "/hello" is accessed but method other than "GET" ex- "POST"
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	// printing the msg
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

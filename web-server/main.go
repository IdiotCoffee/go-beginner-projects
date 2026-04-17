package main

import (
	"fmt"
	"log"
	"net/http"
)

// any API will take a request, and return a response.
// so, pay attention to the parameters that the function takes.
func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello Gophers!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	// again another statement;statement, and another error handling!!!
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	if name == "" || address == "" {
		http.Error(w, "missing form fields", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address - %s\n", address)
	fmt.Fprintf(w, "POST request successful!\n")

}

func main() {
	// point Go to the folder hosting the static files.
	// index.html is automatically picked up.
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer) // the root path takes you to fileserver
	// automatically picks up the index.html
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080...")
	// statement;statement use!!! error handling!!!
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

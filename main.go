package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	fmt.Fprint(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productsHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

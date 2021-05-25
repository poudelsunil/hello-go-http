package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

type server struct{}


func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	if(r.Method == "GET"){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world"}`))
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not Found"}`))
	}
	
}

func main(){
	r := mux.NewRouter()

	// http.Handle("/", s)
	
	// log.Fatal(http.ListenAndServe(":8080",nil))

	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", notFound)

	fmt.Printf("Listining to 8080 port")
    log.Fatal(http.ListenAndServe(":8080", r))

	
}

func get(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET "}`))
}

func post(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "POST "}`))
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "NotFound"}`))
}
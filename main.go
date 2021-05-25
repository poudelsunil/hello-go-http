package main

import (
	"log"
	"net/http"
	"fmt"
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
	s := &server{}
	http.Handle("/", s)
	fmt.Printf("Listining to 8080 port")
	log.Fatal(http.ListenAndServe(":8080",nil))
	
}
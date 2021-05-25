package main

import (
	"log"
	"net/http"
	"fmt"
	"strconv"		// Package strconv implements conversions to and from string representations of basic data types. 

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
	api := r.PathPrefix("/api/v1").Subrouter()

	// http.Handle("/", s)
	
	// log.Fatal(http.ListenAndServe(":8080",nil))

	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)

	api.HandleFunc("/user/{userID}/comment/{commentID}", functionWithParams).Methods(http.MethodGet)

	api.HandleFunc("/", notFound)

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



func functionWithParams(w http.ResponseWriter, r *http.Request) {


    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    userID := -1
    var err error
    if val, ok := pathParams["userID"]; ok {

        userID, err = strconv.Atoi(val)  //  strconv.Atoi (string to int) and strconv.Itoa (int to string). 
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    commentID := -1
    if val, ok := pathParams["commentID"]; ok {
        commentID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    query := r.URL.Query()
    location := query.Get("location")

    w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}
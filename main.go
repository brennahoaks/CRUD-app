package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type foo struct {
	Name string `json:"name"`
	Id   string `json:"id,omitempty"`
}

var fooMap = make(map[string]foo)

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/foo", postFoo)
	r.HandleFunc("/foo/{id}", getFoo)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postFoo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var f foo
		var id = uuid.New().String()
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f.Id = id
		fooMap[id] = f
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(f)
	default:
		fmt.Fprintf(w, "Request method %s is not supported", r.Method)
	}
}

func getFoo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		if val, ok := fooMap[vars["id"]]; !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			json.NewEncoder(w).Encode(val)
		}
	case "DELETE":
		if val, ok := fooMap[vars["id"]]; !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			delete(fooMap, vars["id"])
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(val)
		}
	default:
		fmt.Fprintf(w, "1Request method %s is not supported", r.Method)
	}
}

func main() {
	handleRequests()
}

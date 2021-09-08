package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! I am the ip resolver server!")
}

func CreateRouting() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/", sayHello)
	return r, nil
}

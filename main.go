package main

import (
	"fmt"
	"net/http"

	"github.com/azer/logger"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")

}
func main() {

	var log = logger.New("avoxi-whitelist")
	log.Info("Server has been started")

	r := mux.NewRouter()

	r.HandleFunc("/whitelist", CheckIsWhitelisted).Methods("POST")

	http.ListenAndServe(":8080", r)
}

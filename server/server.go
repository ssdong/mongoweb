package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"mongoweb/command"

	"github.com/gorilla/mux"
)

var commandManager = new(command.CmdManager)

func Listen(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/dbs", getDatabasesHandler)
	// @TODO Need to update the regular expression here to match all possible naming restrictions
	// in Mongo. https://docs.mongodb.com/v3.4/reference/limits/#naming-restrictions
	// Currently, they only match alpahabetic characters
	router.HandleFunc("/dbs/{db:[A-z]+}/collections", getCollectionsHandler)
	router.HandleFunc("/dbs/{db:[A-z]+}/collections/{collection:[A-z]+}", getCollectionHandler)
	log.Fatal(http.ListenAndServe(port, router))
}

func getDatabasesHandler(w http.ResponseWriter, req *http.Request) {
	dbs := commandManager.GetDbs()
	json.NewEncoder(w).Encode(dbs)
}

func getCollectionsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	db := vars["db"]
	collections, err := commandManager.GetCollections(db)
	if err != nil {
		fmt.Fprintln(w, "Error", err)
		return
	}
	json.NewEncoder(w).Encode(collections)
}

func getCollectionHandler(w http.ResponseWriter, req *http.Request) {

}

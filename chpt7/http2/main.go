package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dbServer := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:5000", dbServer))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			msg := fmt.Sprintf("no such item %q\n", item)
			http.Error(w, msg, http.StatusNotFound)
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound)
	}
}

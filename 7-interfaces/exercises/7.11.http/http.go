package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/", db.list)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item required\n")
		return
	}

	price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if price < 0.01 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %v\n", price)
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "%s already exists\n", item)
		return
	}

	db[item] = dollars(price)
	db.list(w, req)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
		return
	}
	fmt.Fprintf(w, "no such item: %s\n", item)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %s\n", item)
		return
	}

	price, _ := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if price < 0.01 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %v\n", price)
		return
	}

	db[item] = dollars(price)
	db.list(w, req)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	delete(db, item)
	db.list(w, req)
}

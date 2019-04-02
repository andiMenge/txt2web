package main

import (
	"log"
	"net/http"
	"flag"
)

func main() {

	var port, file string

	// Configure flags
	flag.StringVar(&port, "port", "8080", "port for the webserver")
	flag.StringVar(&file, "file", "textfile", "name of the textfile to serve")
	
	flag.Parse()

	// root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		http.ServeFile(w, r, file)
	})
	
	// Static fileserver:
	log.Println("Txt2Web")
	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+ port, nil))
}
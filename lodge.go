package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	portPtr := flag.String("port", ":8080", "Port")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("directory is required")
		os.Exit(1)
	}

	directory := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, directory+r.URL.Path)
			return
		}
		http.ServeFile(w, r, directory+"index.html")
	})

	log.Fatal(http.ListenAndServe(*portPtr, nil))
	fmt.Println("Running at port:", *portPtr)
}

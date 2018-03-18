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
	portPtr := flag.String("port", ":8080", "Port to listen on")
	sslPtr := flag.Bool("ssl", false, "Serve content using SSL")
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

	if *sslPtr {
		log.Fatal(http.ListenAndServeTLS(*portPtr, "cert.pem", "key.pem", nil))
	} else {
		log.Fatal(http.ListenAndServe(*portPtr, nil))
	}
	fmt.Println("Running on port:", *portPtr)
}

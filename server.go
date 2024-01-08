package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed statik
var statik embed.FS

func main() {
	statikFS, err := fs.Sub(statik, "statik")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(statikFS)))
	port := ":3000"
	log.Print("KLab Site listening on :3000...")
	runErr := http.ListenAndServe(port, nil)
	if runErr != nil {
		log.Fatal(runErr)
	}
}

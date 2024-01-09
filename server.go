package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

var appKonfig map[string]string

//go:embed statik
var statik embed.FS

func main() {
	appKonfig = createAppKonfig()

	statikFS, err := fs.Sub(statik, "statik")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(statikFS)))
	http.HandleFunc("/konfig", konfigHandler)
	port := ":3000"
	log.Print("KLab Site listening on :3000...")
	runErr := http.ListenAndServe(port, nil)
	if runErr != nil {
		log.Fatal(runErr)
	}
}

func createAppKonfig() map[string]string {
	title := getEnv("TITLE", "KLab Site")

	appKonf := make(map[string]string)
	appKonf["title"] = title
	return appKonf
}

func getEnv(envName, defaultValue string) string {
	envVar, ok := os.LookupEnv(envName)
	if !ok {
		envVar = defaultValue
	}
	return envVar
}

func konfigHandler(w http.ResponseWriter, r *http.Request) {
	if appKonfig == nil {
		appKonfig = createAppKonfig()
	}
	konfig, err := json.Marshal(appKonfig)
	if err != nil {
		fmt.Fprint(w, "Application Error")
		log.Printf("Failed to Marshal (convert to JSON) appKonfig. Error is: %s \n", err)
		return
	}
	fmt.Fprint(w, string(konfig))
}

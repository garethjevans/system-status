package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/newrelic/go-agent"
)
// Handlers
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {

    nrLicense := os.Getenv("NEW_RELIC_LICENSE_KEY")
	nrName := os.Getenv("NEW_RELIC_APP_NAME")
	port := os.Getenv("PORT")

	config := newrelic.NewConfig(nrName, nrLicense)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Fatal("Unable to create new relic app: ", err)
		os.Exit(1)
	}

	log.Println("Starting application server on port " + port)

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/status", StatusHandler))
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", StatusHandler))

	err2 := http.ListenAndServe(":"+port, nil)
	if err2 != nil {
		log.Fatal("ListenAndServe: ", err2)
		os.Exit(1)
	}
}

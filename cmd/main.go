package main

import (
	"flag"
	"log"
	"net/http"

	restapi "github.com/klebervirgilio/vue-crud-app-with-golang/pkg/http"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/storage"
)

func main() {
	var httpPort string
	flag.StringVar(&httpPort, "b", ":4444", "bind on port")
	flag.Parse()

	repo := storage.NewMongoRepository()
	api := restapi.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, api.Router))
}

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
	handler := restapi.NewService(repo)

	log.Fatal(http.ListenAndServe(httpPort, handler.Router))
}

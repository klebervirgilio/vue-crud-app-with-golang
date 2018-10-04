package main

import (
	"log"
	"net/http"
	"os"

	web "github.com/klebervirgilio/vue-crud-app-with-golang/pkg/http"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/storage"
)

func main() {
	httpPort := os.Getenv("PORT")
	// flag.StringVar(&httpPort, "b", ":4444", "bind on port")
	// flag.Parse()

	repo := storage.NewMongoRepository()
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}

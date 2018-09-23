package main

import (
	"flag"
	"fmt"

	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/core"
	"github.com/klebervirgilio/vue-crud-app-with-golang/pkg/storage"
)

func main() {
	var httpPort string
	flag.StringVar(&httpPort, "bind on port", ":5000", "")
	flag.Parse()

	repo := storage.NewMongoRepository()
	repo.Create(&core.Kudo{})

	fmt.Println("spin up server...")
}

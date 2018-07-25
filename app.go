package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maxsakharov/gatlingcomparison/front/routes"
	"github.com/maxsakharov/gatlingcomparison/storage"
)

func main() {
	host := os.Getenv("HOST")
	if len(host) < 1 {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8089"
	}

	storage := storage.FileStorage{os.Getenv("REPORTS_PATH")}
	addr := fmt.Sprintf("%s:%s", host, port)

	log.Print("Starting server at " + addr + " ...")
	log.Fatal(http.ListenAndServe(addr, routes.InitRoutes(storage)))
}

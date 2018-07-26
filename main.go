package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/maxsakharov/gatlingcomparison/web"
)

func main() {
	conf := Config{}
	web := web.NewWeb(conf.Storage())

	addr := fmt.Sprintf("%s:%s", conf.HttpHost(), strconv.Itoa(conf.HttpPort()))

	log.Print("Starting server at " + addr + " ...")
	log.Fatal(http.ListenAndServe(addr, web.Router))
}

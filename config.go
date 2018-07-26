package main

import (
	"log"
	"os"
	"strconv"

	"github.com/maxsakharov/gatlingcomparison/cmd/gc/storage"
)

type Config struct{}

func (c *Config) HttpPort() int {
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = "8089"
	}

	p, err := strconv.Atoi(port)

	if err != nil {
		log.Panic(err)
	}

	return p
}

func (c *Config) HttpHost() string {
	host := os.Getenv("HOST")
	if len(host) < 1 {
		host = "0.0.0.0"
	}
	return host
}

func (c *Config) ReportsPath() string {
	path := os.Getenv("REPORTS_PATH")
	if len(path) < 1 {
		path = "reports/"
	}
	return path
}

func (c *Config) Storage() storage.Storage {
	return storage.FileStorage{c.ReportsPath()}
}

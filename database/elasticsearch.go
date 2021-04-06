package database

import (
	"fmt"

	"github.com/elastic/go-elasticsearch"
)

func GetESClient() *elasticsearch.Client {
	// Might be redundant
	ES, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}
	return ES
}

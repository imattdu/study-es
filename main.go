package main

import (
	"log"

	// 注意版本：需要和连接的 es 保持一致
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	//es, _ := elasticsearch.NewDefaultClient()
	//log.Println(elasticsearch.Version)
	//log.Println(es.Info())

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"strings"
)

type Hit struct {
	Source struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Age  int64  `json:"age"`
		// omitempty 如果不设置这个字段就不会输出
		// 不加omitempty，如果不设置这个字段也会输出默认值
		H string `json:"h,omitempty"`
	} `json:"_source"`
}

type Hits struct {
	Hits []Hit `json:"hits"`
}

type EsCommonRes struct {
	Took int64 `json:"took"`
	Hits Hits  `json:"hits"`
}

func main() {

	// 1.创建 es 客户端
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Printf("%v", err)
	}
	//var r = make(map[string]interface{})
	var r = EsCommonRes{}

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "matt",
			},
		},
	}

	query = map[string]interface{}{
		"query": map[string]interface{}{
			"fuzzy": map[string]interface{}{
				"name": map[string]interface{}{
					"value": "matt",
				},
			},
		},
		"sort": []map[string]interface{}{
			{
				"age": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"_score": map[string]interface{}{
					"order": "desc",
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("user"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.

	log.Printf("%+v", r)

	//json.Unmarshal(res.Body, &r)

	log.Println(strings.Repeat("=", 37))

}

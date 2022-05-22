package main

// 插入操作
import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	// omitempty 如果不设置这个字段就不会输出
	// 不加omitempty，如果不设置这个字段也会输出默认值
	H string `json:"h,omitempty"`
}

func main() {

	// 1.创建 es 客户端
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Printf("%v", err)
	}

	// 2.请求的数据
	data, err := json.Marshal(User{Id: 2, Name: "matt", Age: 17, H: "hh"})
	if err != nil {
		log.Fatalf("Error marshaling document: %s", err)
	}

	// 3.创建请求对象
	req := esapi.IndexRequest{
		Index: "user",
		// 可以不指定id
		// DocumentID: "1",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	// 4.发送请求
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {

	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			log.Printf("%+v", r)
		}
	}

}

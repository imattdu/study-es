package main

import (
	"encoding/json"
	"log"
)

type User1 struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	H    string `json:"h,omitempty"` // omitempty 忽略
}

func main() {

	var u1 = User1{
		Id:  1,
		Age: 13,
		//Name: "matt",
		//H: "h1",
	}

	m1, _ := json.Marshal(u1)
	log.Println(string(m1))
}

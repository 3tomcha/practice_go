package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json: "user_id"`
	UserName  string   `json: "user_name"`
	Languages []string `json: "language"`
}

func main() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Record struct {
	ProcessID string `json: "process_id"`
	DeletedAt JSTime `json: "deleted_at"`
}

type JSTime time.Time

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

func main() {
	r := &Record{
		ProcessID: "0001",
		DeletedAt: JSTime(time.Now()),
	}
	b, _ := json.MarshalJSON(&r)
	fmt.Println(string(b))
}

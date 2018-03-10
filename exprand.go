package main

import (
	"fmt"
	"math/rand"
    "time"
    "net/http"
)

func main() {
	for {
		time_curl := rand.ExpFloat64() / 1
        time_curl = time_curl * 10
        sp := int64(time_curl)
        fmt.Printf("%ds, send curl to localhost:30001/stress\n", sp)
        time.Sleep(time.Duration(sp) * time.Second)
		resp, err := http.Get("http://localhost:30001/stress")
		if err != nil {
            panic(err)
		}
		defer resp.Body.Close()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	for {
		time_curl := rand.ExpFloat64() / 5
		time_curl = time_curl * 10
		time.Sleep(time.Duration(int(time_curl)) * time.Second)
		sp := strconv.Itoa(int(time_curl))
		fmt.Printf("%s s, send curl to localhost:30001/stress\n", sp)
		cmd := exec.Command("bash", "curl2.sh", sp)
		cmd.Run()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("please input 3 args, ex: /bin [lambda]")
	}
	lambda, _ := strconv.Atoi(os.Args[1])
	for {
		time_curl := rand.ExpFloat64() / float64(lambda)
		time_curl = time_curl * 10
		time.Sleep(time.Duration(int(time_curl)) * time.Second)
		sp := strconv.Itoa(int(time_curl))
		fmt.Printf("%s s, send curl to localhost:30001/stress\n", sp)
		cmd := exec.Command("bash", "curl2.sh", sp)
		cmd.Run()
	}
}

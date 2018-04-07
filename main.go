package main

import (
	"github.com/gin-gonic/gin"
	"os/exec"
)

func test() {
	cmd := exec.Command("stress", "-c", "3", "-m", "1", "--vm-bytes", "200M", "-t", "15s")
	cmd.Run()
}

func stress(c *gin.Context) {
	go test()
}

func main() {
	r := gin.Default()
	r.GET("/stress", stress)
	r.Run(":8888")
}

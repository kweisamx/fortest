package main

import(
    "os/exec"
    "github.com/gin-gonic/gin"
)

func test(){
    cmd:= exec.Command("stress","-i","1","-t","15s")
    cmd.Run()
}

func stress(c *gin.Context){
    go test()
}

func main(){
    r := gin.Default()
    r.GET("/stress", stress)
    r.Run(":8888")
}

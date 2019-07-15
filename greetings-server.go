package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	Version  string
	Revision string
)

func executeIt() string {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("in all caps: %q\n", out.String())
	return out.String()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Global middleware
	router.Use(gin.Logger())

	// Routing
	router.StaticFile("/", "./index.html")

	router.GET("/api/exec", func(ctx *gin.Context) {
		res := executeIt()

		ctx.JSON(200, gin.H{
			"result": res,
		})
	})

	return router
}

func main() {

	fmt.Println("Greetings Server : Version:" + Version + " Revision:" + Revision)

	endless.ListenAndServe(":8080", setupRouter())
}

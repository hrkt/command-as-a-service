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
	// Version number
	Version string
	// Revision number
	Revision string
	// RequestBodyBufferSize is the buffer size for http-request-body
	RequestBodyBufferSize = 2048
)

func executeIt(requestBody string) string {
	cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("some input")
	cmd.Stdin = strings.NewReader(requestBody)
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

	router.POST("/api/exec", func(ctx *gin.Context) {
		buf := make([]byte, RequestBodyBufferSize)
		n, _ := ctx.Request.Body.Read(buf)
		body := string(buf[0:n])

		res := executeIt(body)

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

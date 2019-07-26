package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/fvbock/endless"
)

// AppConfig is a struct for app-config.json
type AppConfig struct {
	// a command to execute
	Command string `json:"command"`
	// arguments for the command
	Arguments []string `json:"arguments"`
}

var (
	// Version number
	Version string
	// Revision number
	Revision string
	// RequestBodyBufferSize is the buffer size for http-request-body
	RequestBodyBufferSize = 2048
	// application configuration
	appConfig AppConfig
)

func init() {
	file, err := ioutil.ReadFile("app-settings.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &appConfig)
	fmt.Printf("Command :%s\n", appConfig.Command)
}

func executeIt(requestBody string) string {
	cmd := exec.Command(appConfig.Command, appConfig.Arguments[:]...)
	cmd.Stdin = strings.NewReader(requestBody)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func main() {

	fmt.Println("command-as-a-service : Version:" + Version + " Revision:" + Revision)

	endless.ListenAndServe(":8080", setupRouter())
}

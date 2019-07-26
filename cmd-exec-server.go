package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fvbock/endless"
)

// AppConfig is a struct for app-config.json
type AppConfig struct {
	// a command to execute
	Port int `json:"port"`
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
	fmt.Printf("Port :%d\n", appConfig.Port)
}

func executeIt(path string, requestBody string) string {
	//cmd := exec.Command(path, appConfig.Arguments[:]...)
	cmd := exec.Command(path)
	cmd.Stdin = strings.NewReader(requestBody)
	out, err := cmd.Output()
	if err != nil {
		log.Printf(err.Error())
		return string(err.Error())
	}
	return string(out)
}

func MyServer() http.Handler {
	return &myHandler{}
}

type myHandler struct {
}

func (f *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("called:" + r.URL.Path)
	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf(err.Error())
	}
	res := executeIt(r.URL.Path, string(buffer))
	w.Write([]byte(res))
}

func main() {

	fmt.Println("command-as-a-service : Version:" + Version + " Revision:" + Revision)

	myHandler := MyServer()

	endless.ListenAndServe(":"+strconv.Itoa(appConfig.Port), myHandler)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func executeIt(path string, requestBody string, params []string) string {
	var cmd *exec.Cmd
	if len(params) > 0 && params[0] == "" {
		cmd = exec.Command(path)
	} else {
		cmd = exec.Command(path, params[:]...)
	}
	cmd.Stdin = strings.NewReader(requestBody)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Printf(err.Error())
		return string(err.Error())
	}
	log.Println(stdout.String())
	return string(stdout.String())
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

	unescaped, _ := url.QueryUnescape(r.URL.RawQuery)
	params := strings.Split(unescaped, "&")
	log.Println(params)
	res := executeIt(r.URL.Path, string(buffer), params)
	w.Write([]byte(res))

}

func main() {

	fmt.Println("command-as-a-service : Version:" + Version + " Revision:" + Revision)

	myHandler := MyServer()

	endless.ListenAndServe(":"+strconv.Itoa(appConfig.Port), myHandler)
}

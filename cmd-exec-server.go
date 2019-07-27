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
	// if true, do not use whitelist
	DangerousMode bool `json:"dangerousMode"`
	// a command to execute
	Port int `json:"port"`
	// arguments for the command
	Whitelist []string `json:"whitelist"`
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
	// whitelistMap holds entries of whitelist as map
	whitelistMap map[string]struct{}
)

func init() {
	file, err := ioutil.ReadFile("app-settings.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &appConfig)
	whitelistMap = make(map[string]struct{})
	for _, v := range appConfig.Whitelist {
		whitelistMap[v] = struct{}{}
	}

	fmt.Printf("Port :%d\n", appConfig.Port)
}

func executeIt(path string, requestBody string, params []string) (int, string) {

	if !appConfig.DangerousMode {
		_, ok := whitelistMap[path]
		if !ok {
			return 400, string("ERROR: command " + path + " not in the whitelist")
		}
	}

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
		return 400, string("ERROR: command execution failed. reason: " + err.Error())
	}
	log.Println(stdout.String())
	return 0, string(stdout.String())
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
	code, res := executeIt(r.URL.Path, string(buffer), params)
	if code == 400 {
		http.Error(w, res, http.StatusBadRequest)
		return
	}
	w.Write([]byte(res))
}

func main() {

	fmt.Println("command-as-a-service : Version:" + Version + " Revision:" + Revision)

	myHandler := MyServer()

	endless.ListenAndServe(":"+strconv.Itoa(appConfig.Port), myHandler)
}

package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8000", nil)
}

var routes = map[string]func(response http.ResponseWriter, request *http.Request){
	"/hello":       hello,
	"/file":        file,
	"/sleep/10ms":  sleep10ms,
	"/sleep/100ms": sleep100ms,
	"/sleep/1s":    sleep1s,
	"/sleep/5s":    sleep5s,
}

func handler(response http.ResponseWriter, request *http.Request) {
	routes[request.URL.Path](response, request)
}

func hello(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello"))
}

func file(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "file.txt")
}

func sleep10ms(response http.ResponseWriter, request *http.Request) {
	time.Sleep(10 * time.Millisecond)
	response.Write([]byte("ok"))
}

func sleep100ms(response http.ResponseWriter, request *http.Request) {
	time.Sleep(100 * time.Millisecond)
	response.Write([]byte("ok"))
}

func sleep1s(response http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Second)
	response.Write([]byte("ok"))
}

func sleep5s(response http.ResponseWriter, request *http.Request) {
	time.Sleep(5 * time.Second)
	response.Write([]byte("ok"))
}

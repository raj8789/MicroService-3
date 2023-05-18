package main

import (
	"log"
	//handle "micro/handler"
	"github.com/raj8789/MicroService-3/handle"
	"net/http"
	"os"
)

func main() {
	sm := http.NewServeMux()
	l := log.New(os.Stdout, "Example-API", log.LstdFlags)
	h := handle.NewHello(l)
	sm.Handle("/more", h)
	http.ListenAndServe(":8080", sm)
}

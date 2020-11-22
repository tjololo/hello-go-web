package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var port int
	var responseString string
	flag.IntVar(&port, "port", 8080, "Port the webserver should launch with")
	flag.StringVar(&responseString, "response", "Hello from go code", "Content of respone element")
	flag.Parse()
	log.Printf("Staring application on port %d...\n", port)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request")
		response := WebResponse{
			Response: responseString,
			Info:     fmt.Sprintf("Application info; %s", os.Getenv("APP_INFO")),
		}
		responseBytes, _ := json.Marshal(response)
		log.Println(fmt.Sprintf("%v", response))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(responseBytes)
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	server := http.Server{Addr: fmt.Sprintf(":%d", port)}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown received, exiting...")

	server.Shutdown(context.Background())
}

//WebResponse struct returned from webendpoint
type WebResponse struct {
	Response string `json:"response"`
	Info     string `json:"info"`
}

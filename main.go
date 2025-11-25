package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
)

var responseString string

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port the webserver should launch with")
	flag.StringVar(&responseString, "response", "Hello from go code", "Content of respone element")
	flag.Parse()
	log.Printf("Staring application on port %d...\n", port)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/errors", errorHandler)

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

	err := server.Shutdown(context.Background())

	if err != nil {
		log.Fatal(err)
	}
}

// WebResponse struct returned from webendpoint
type WebResponse struct {
	Response string `json:"response"`
	Info     string `json:"info"`
}

// SortWebResponse by Info
func SortWebResponseByInfo(responses []WebResponse) []WebResponse {
	sort.Slice(responses, func(i, j int) bool {
		return responses[i].Info < responses[j].Info
	})
	return responses
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got a request")
	response := WebResponse{
		Response: responseString,
		Info:     fmt.Sprintf("Application info; %s", os.Getenv("APP_INFO")),
	}
	responseBytes, _ := json.Marshal(response)
	log.Printf("%v\n", response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	_, err := w.Write(responseBytes)
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	serverErrorRate := readIntParam("500", r)
	nBig, err := rand.Int(rand.Reader, big.NewInt(101))
	if err != nil {
		http.Error(w, "Failed to generate random number", http.StatusInternalServerError)
		return
	}
	rNumber := int(nBig.Int64())
	if rNumber < serverErrorRate {
		http.Error(w, fmt.Sprintf("500 error returned as %d < %d", rNumber, serverErrorRate), http.StatusInternalServerError)
	} else {
		_, err = w.Write([]byte(fmt.Sprintf("OK as %d >= %d", rNumber, serverErrorRate)))
		if err != nil {
			log.Printf("Error writing response: %v\n", err)
		}
	}
}

func readIntParam(name string, r *http.Request) (value int) {
	serverErrorRate := r.URL.Query().Get(name)
	if i, err := strconv.Atoi(serverErrorRate); err == nil {
		value = i
	}
	return
}

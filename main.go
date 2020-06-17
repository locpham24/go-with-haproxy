package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func createServer(port int) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world at "+strconv.Itoa(port))
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}
	return &server
}

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		server := createServer(9000)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	go func() {
		server := createServer(9001)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	go func() {
		server := createServer(9002)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}

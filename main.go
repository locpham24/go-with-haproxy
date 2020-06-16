package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main(){
	wg := new(sync.WaitGroup)

	wg.Add(3)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world at " + r.Host)
	})

	go func(){
		log.Fatal(http.ListenAndServe(":9000", nil))
		wg.Done()
	}()

	go func(){
		log.Fatal(http.ListenAndServe(":9001", nil))
		wg.Done()
	}()

	go func(){
		log.Fatal(http.ListenAndServe(":9002", nil))
		wg.Done()
	}()

	wg.Wait()
}
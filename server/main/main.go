package main

import (
	"net/http"
	"sinergy-test/server"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		http.ListenAndServe(":8000", server.ListenData())
		wg.Done()
	}()
	wg.Wait()
}

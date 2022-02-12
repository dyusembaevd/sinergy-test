package main

import (
	"net/http"
	"sinergy-test/client"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	client.StartCron()
	go func() {
		http.ListenAndServe(":8001", client.GetAction())
		wg.Done()
	}()
	wg.Wait()
}

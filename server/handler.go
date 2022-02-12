package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://novasite.su/test1.php",
	"https://novasite.su/test2.php",
}

func ListenData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Transfer-Encoding", "chunked")
		w.WriteHeader(http.StatusOK)

		ch := make(chan []byte, len(urls))
		for {
			GetAllData(ch)

			cn, ok := w.(http.CloseNotifier)
			if !ok {
				http.NotFound(w, r)
				return
			}
			flusher, ok := w.(http.Flusher)
			if !ok {
				http.NotFound(w, r)
				return
			}

			flusher.Flush()

			enc := json.NewEncoder(w)

			for len(ch) != 0 {
				select {
				case <-cn.CloseNotify():
					return
				case data := <-ch:
					if len(data) == 0 {
						continue
					}
					fmt.Println(string(data))
					err := enc.Encode(data)
					if err != nil {
						return
					}
					flusher.Flush()
				}
			}
			<-time.After(1 * time.Second)
		}
	}
}

func GetAllData(ch chan []byte) {
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			var data []byte
			var err error

			if data, err = GetData(url); err != nil {
				fmt.Printf("Error: %")
			}

			ch <- data

			wg.Done()
		}(url)
	}
	wg.Wait()
}

func GetData(url string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

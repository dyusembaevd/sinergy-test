package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type SrvResponse struct {
	Action *string      `json:"action"`
	Type   *string      `json:"type"`
	Data   *interface{} `json:"data"`
}

var Buff []byte

func StartCron() {
	for {
		req, err := http.NewRequest("GET", "http://localhost:8000/listen-data", nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return
		}

		dec := json.NewDecoder(resp.Body)
		for {
			err := dec.Decode(&Buff)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			fmt.Printf("Got response: %+v\n", string(Buff))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

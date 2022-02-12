package client

import (
	"encoding/json"
	"net/http"
	"time"
)

func GetAction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp SrvResponse
		var err error
		timer := time.NewTimer(1 * time.Second)

		for resp.Action == nil {
			err = json.Unmarshal(Buff, &resp)
		}

		if err != nil {
			jsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"error": err,
			})
			return
		}

		select {
		case <-timer.C:
			jsonResponse(w, http.StatusOK, "TIMEOUT")
		}
		jsonResponse(w, http.StatusOK, resp)
	}
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

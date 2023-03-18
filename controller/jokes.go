package controller

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Joke struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	Value   string `json:"value"`
	IconUrl string `json:"icon_url"`
}

type BaseResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func response(w http.ResponseWriter, data interface{}, message string, status int, errStatus bool) {
	var response BaseResponse
	response.Status = "success"
	if errStatus {
		response.Status = "error"
	}

	if data != nil {
		response.Data = data
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func GetJokes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response(w, nil, "Method not allowed", http.StatusBadRequest, true)
		return
	}

	jokes := []Joke{}
	makeRequest(25, &jokes)

	response(w, jokes, "Success", http.StatusOK, false)
}

func makeRequest(missingItems int, jokes *[]Joke) error {
	lock := sync.RWMutex{}
	done := make(chan struct{})

	for i := 0; i < missingItems; i++ {
		go func() {
			resp, err := http.Get("https://api.chucknorris.io/jokes/random")
			if err != nil {
				return
			}

			var joke Joke
			err = json.NewDecoder(resp.Body).Decode(&joke)
			if err != nil {
				return
			}

			lock.RLock()
			for _, jok := range *jokes {
				if jok.Id == joke.Id {
					continue
				}
			}
			lock.RUnlock()

			lock.Lock()
			defer lock.Unlock()
			*jokes = append(*jokes, joke)
			if len(*jokes) == 25 {
				done <- struct{}{}
			}
		}()
	}
	<-done

	if len(*jokes) < 25 {
		makeRequest(25-len(*jokes), jokes)
	}
	return nil
}

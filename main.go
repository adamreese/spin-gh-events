package main

import (
	"encoding/json"
	"log"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/fermyon/spin/sdk/go/key_value"
	"github.com/google/go-github/v51/github"
)

func getSecretKey() ([]byte, error) {
	store, err := key_value.Open("default")
	if err != nil {
		return nil, err
	}
	defer key_value.Close(store)

	key, err := key_value.Get(store, "webhook_secret")
	return key, err
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		webhookSecretKey, err := getSecretKey()
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload, err := github.ValidatePayload(r, webhookSecretKey)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		webHookType := github.WebHookType(r)
		log.Printf("Event received %q", webHookType)

		switch webHookType {
		case "ping":
			log.Print("ping")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		case "push":
			var event github.PushEvent
			if err := json.Unmarshal(payload, &event); err != nil {
				log.Print(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			// log the request
			log.Printf("%v", event)

			// add any other events to handle
		}
	})
}

func main() {}

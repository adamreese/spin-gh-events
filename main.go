package main

import (
	"encoding/json"
	"log"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/fermyon/spin/sdk/go/key_value"
	"github.com/google/go-github/v51/github"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		store, err := key_value.Open("default")
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer key_value.Close(store)

		webhookSecretKey, err := key_value.Get(store, "webhookSecretKey")
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

		switch github.WebHookType(r) {
		case "ping":
			log.Print("ping")
			w.WriteHeader(http.StatusOK)
		case "push":
			var event github.PushEvent
			if err := json.Unmarshal(payload, &event); err != nil {
				log.Print(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("%v", event)
		}
	})
}

func main() {}

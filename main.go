package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func snsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only POST accepted", http.StatusMethodNotAllowed)
		return
	}

	if strings.HasSuffix(r.URL.Path, "/with-auth") {
		u, p, ok := r.BasicAuth()
		if !ok {
			log.Printf("at=auth ok=false")
			w.Header().Set("WWW-Authenticate", "Basic realm=\"hello\"")
			http.Error(w, "unauth", http.StatusUnauthorized)
			return
		}
		log.Printf("at=auth ok=true", u, p)

		if !(u == "user" && p == "pass") {
			log.Printf("at=auth ok=true creds=bad", u, p)
			http.Error(w, "unauth", http.StatusUnauthorized)
			return
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("at=read-body err=%q", err)
		http.Error(w, "read body failed", http.StatusInternalServerError)
		return
	}

	var subReq struct {
		SubscribeURL string
	}

	if err := json.Unmarshal(body, &subReq); err != nil {
		log.Printf("at=decode err=%q", err)
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}

	if subReq.SubscribeURL != "" {
		resp, err := http.Get(subReq.SubscribeURL)
		if err != nil {
			log.Printf("at=hit-subscribe-error err=%q")
			http.Error(w, "subscribe failed", http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Printf("at=hit-subscribe-error err=%q", fmt.Sprintf("wanted status 200, got %d", resp.StatusCode))
			http.Error(w, "subscribe failed", http.StatusBadRequest)
			return
		}
		log.Printf("at=hit-subscribe-success")
		return
	}

	log.Printf("at=notification")
	log.Printf("%s", string(body))
}

func main() {
	http.HandleFunc("/sns", snsHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func snsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("at=snsHandler request=%+v\n", r)
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

	if forward != nil {
		// parse body
		type alarmReq struct {
			Message string
		}
		var data alarmReq

		err := json.Unmarshal(body, &data)
		if err != nil {
			log.Printf("at=forward:unmarshal error=%v\n", err)
			return
		}

		var message []byte
		message = []byte(data.Message)
		log.Printf("at=forward:message message=%s\n", message)

		req, err := http.NewRequest("POST", *forward, bytes.NewBuffer(message))
		if err != nil {
			log.Printf("at=forward:new_request error=%v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("at=forward:post error=%v\n", err)
			return
		}
		defer resp.Body.Close()

		log.Printf("at=forward:response status=%s\n", resp.Status)
	}
}

var forward *string

func main() {
	if f := os.Getenv("FORWARD"); f != "" {
		// validate url
		_, err := url.Parse(f)
		if err != nil {
			log.Fatal("at=main:forward error=%v\n", err)
		}

		forward = &f
		log.Println("at=main forward=" + *forward)
	}

	http.HandleFunc("/sns", snsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := os.Getenv("ADDR")

	log.Printf("at=main:startup listener=%s:%s", addr, port)
	log.Fatal(http.ListenAndServe(addr+":"+port, nil))

}

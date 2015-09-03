package main

import (
	"github.com/jmervine/sns-fwd/confirmation"
	"github.com/jmervine/sns-fwd/forwarder"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jmervine/sns-fwd/Godeps/_workspace/src/github.com/jmervine/env"
	log "github.com/jmervine/sns-fwd/Godeps/_workspace/src/github.com/jmervine/readable"
)

func handler(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	status := http.StatusOK // default to 200, override when error

	// always log method, path, status and took when finishing
	defer func(code int) {
		log.Print(
			"at=main on=handler method", r.Method,
			"path", r.URL.Path,
			"status", code,
			"took", time.Since(begin),
		)
	}(status)

	// handler http.Error and set status for use in above
	// defer function
	httpError := func(code int) {
		status = code
		http.Error(w, http.StatusText(code), code)
	}

	if r.Method != "POST" {
		httpError(http.StatusMethodNotAllowed)
		return
	}

	// handle basic auth when USERNAME and PASSWORD are set
	username := env.Get("USERNAME")
	password := env.Get("PASSWORD")

	if username != "" && password != "" {
		user, pass, ok := r.BasicAuth()

		if !ok {
			log.Print("at=main on=handler action=<req>.BasicAuth() ok=false")

			// support AWS's two attempt authentication
			w.Header().Set("WWW-Authenticate", "Basic realm=\"hello\"")

			httpError(http.StatusUnauthorized)
			return
		}

		log.Debug("at=main on=handler action=<req>.BasicAuth() ok=true user", user, "pass", pass)

		if !(user == username && pass == password) {
			log.Print("at=main on=handler action=<req>.BasicAuth() error", http.StatusText(http.StatusUnauthorized))

			httpError(http.StatusUnauthorized)
			return
		}
	}

	// parse body and deal
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("at=main on=handler action=ioutil.ReadAll(body) error", err)

		httpError(http.StatusInternalServerError)
		return
	}

	log.Debug("at=main on=handler body", body)

	confirm := confirmation.Parse(body)
	if confirm != nil {
		err = confirm.Do()
		if err != nil {
			log.Print("at=main on=handler action=<Confirmation>.Do() error", err)

			httpError(http.StatusBadRequest)
			return
		}

		log.Print("at=main on=handler action=<Confirmation>.Do() ok=true")
		return
	}

	fwdr := forwarder.New(body)
	if err := fwdr.Do(); err != nil {
		log.Print("at=main on=handler action=<Forwarder>.Do() error", err)
		httpError(http.StatusRequestedRangeNotSatisfiable)
	}

	log.Print("at=main on=handler action=<Forwarder>.Do() ok=true")
}

func main() {
	// load '.env' file if present
	env.Load()

	// ensure FORWARD_URL
	if _, err := env.Require("FORWARD_URL"); err != nil {
		log.Fatal("at=main on=env.Require error", err)
	}

	http.HandleFunc("/", handler)

	listener := fmt.Sprintf(
		"%s:%s",
		env.Get("ADDR"),
		env.GetOrSet("PORT", "3000"),
	)

	log.Print("at=main on=startup listener", listener)
	log.Fatal(http.ListenAndServe(listener, nil))
}

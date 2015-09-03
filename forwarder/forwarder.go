package forwarder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmervine/env"
	log "github.com/jmervine/readable"
)

type Forwarder struct {
	dest string // destination url
	body []byte // post body
}

func New(body []byte) *Forwarder {
	f := new(Forwarder)
	f.dest = env.Get("FORWARD_URL")

	// used to parse alarm message and determine if valid json for other fowards
	var data struct {
		Message string
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Debug("at=ParseForwarder on=json.Unmarshal error", err)
		return nil
	}

	f.body = body

	// override body w/ alarm message
	if env.GetBool("FORWARD_ALARM") {
		f.body = []byte(data.Message)
	}

	return f
}

func (f *Forwarder) Do() error {
	resp, err := http.Post(f.dest, "application/json", bytes.NewBuffer(f.body))
	if err != nil {
		log.Debug("at=forwarder.Do() on=http.Post error", err)
		return err
	}

	if resp.StatusCode > 299 {
		err = fmt.Errorf("unexecpted response status %s", err)
		log.Debug("at=fowarder on=Do error", err)
		return err
	}

	return nil
}

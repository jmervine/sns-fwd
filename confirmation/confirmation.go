package confirmation

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/jmervine/readable"
)

type Confirmation struct {
	url string `json:"SubscribeURL"`
}

func Parse(body []byte) *Confirmation {
	c := new(Confirmation)

	if err := json.Unmarshal(body, c); err != nil {
		log.Debug("at=ParseConfirmation err", err)
		return nil
	}

	if c.url == "" {
		return nil
	}

	return c
}

func (c *Confirmation) Do() error {
	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return fmt.Errorf("Confirmation failed with code: %d", resp.StatusCode)
	}

	return nil
}

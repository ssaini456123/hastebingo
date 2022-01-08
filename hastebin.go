package hastebingo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var base = "https://hastebin.com/documents"

// Hastebin holds the basic structure of
// hastebin hastes
//
// Please note however that this is based on my observations
// from the element inspector in my browser.
type Hastebin struct {
	// Key is the unique key generated for each haste.
	//
	// This unique key is usually found at the end of
	// the url once a new haste is created.
	Key string `json:"key"`
}

// Post posts a hastebin document and returns the key to that haste.
// The haste key returned will contain exactly what you put in as
// the contents.
//
// The contents of this haste in our case is the `data` parameter
//
func (h *Hastebin) Post(data string) (string, error) {
	// Creates a json body.
	rbody := []byte(data)

	// Create a new post request and creates a new bytebuffer
	// with our given Data.
	post, err := http.Post(base, "application/json", bytes.NewBuffer(rbody))
	if err != nil {
		return "", err

	}
	// Time to decode our body to defer the key.
	hb := Hastebin{}
	if err := json.NewDecoder(post.Body).Decode(&hb); err != nil {
		return "", err
	}

	// Return our key.
	return hb.Key, nil
}

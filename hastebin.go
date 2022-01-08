package hastebingo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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

	Data string `json: "data"`
}

var key string

// Post posts a hastebin document to the hastebin server
// The contents of this haste in our case is the `data` parameter
//
// To retrieve the key of the created haste, see the RetrieveKey() method.
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

	key = hb.Key

	return "", err
}

// Reads the contents of a hastebin haste.
// Takes in the key to the document and returns the raw string of the haste.
//
// To find the key to the haste you created, see RetrieveKey()
func (h *Hastebin) Read(key string) (string, error) {
	url := base + "/"
	resp, err := http.Get(url + key)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	haste := Hastebin{}

	jsonError := json.Unmarshal(body, &haste)

	if jsonError != nil {
		log.Fatalln(err)
		return "", err
	}

	return haste.Data, err
}

// RetrieveKey returns the key of the haste created from Post().
// This key can be used to access a hastebin document or read from a hastebin
// document using the Read() method.
func (h *Hastebin) RetrieveKey() string {
	return key
}

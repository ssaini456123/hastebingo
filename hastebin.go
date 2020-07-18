package hastebingo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

var base = "https://hastebin.com/"

// Hastebin holds the basic structure of
// hastebin hastes
//
// Please note however that this is based on my observations
// from the element inspector in my browser.
type Hastebin struct {
	// Key is the key generated for the haste.
	//
	// It is usually found at the end of the url once
	// a haste is created.
	//
	// I recommend you only fill this field in when you want
	// to get information from a specific haste instead of
	// creating a haste.
	Key string

	// Data is the contents of the haste.
	Data string
}

// Post posts a hastebin document and returns the key to that haste,
// The haste key returned will contain exactly what you put in as
// the contents.
func (h *Hastebin) Post() string {
	// Creates a json body.
	rbody := []byte(h.Data)

	// Create a new post request and creates a new bytebuffer
	// with our given Data.
	post, err := http.Post(base, "application/json", bytes.NewBuffer(rbody))

	if err != nil {
		// Something went wrong, panic.
		panic(err)
	}

	// Read our request body
	body, err := ioutil.ReadAll(post.Body)

	// Return the request body casted string.
	return string(body)
}

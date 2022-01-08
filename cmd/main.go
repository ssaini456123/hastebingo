package main

import (
	"fmt"

	"github.com/sa111n111/hastebingo"
)

func main() {
	h := hastebingo.Hastebin{}
	h.Post("HI EVERYONE")
	key := h.RetrieveKey()

	fmt.Println(key)

	fmt.Println(h.Read(key)) // read from key
}

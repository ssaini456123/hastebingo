package main

import (
	"fmt"

	"github.com/sa111n111/hastebingo"
)

func main() {
	h := hastebingo.Hastebin{}
	h.PasteFile("to_be_sent.txt")
	key := h.GetKey()

	fmt.Println("The key: " + key)
	result, err := h.Read(key) // read from key.

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

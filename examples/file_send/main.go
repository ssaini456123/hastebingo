package main

import (
	"fmt"

	"github.com/sa111n111/hastebingo"
)

func main() {
	h := hastebingo.Hastebin{}
	h.PasteFile("to_be_sent.txt")
	key := h.RetrieveKey()

	fmt.Println(key)

	fmt.Println(h.Read(key)) // read from key
}

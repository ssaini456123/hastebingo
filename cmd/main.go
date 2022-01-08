package main

import (
	"fmt"

	"github.com/sa111n111/hastebingo"
)

func main() {
	h := hastebingo.Hastebin{}
	fmt.Println(h.Post("HI EVERYONE!"))
}

# hastebingo
---

⚠️ ` Due to craiyon changing the way the requests to their API works, as well as zero development in the past years. Hastebingo is no longer supported in modern version.
Please fork the repository if you would like to continue its development.`

---


A small Hastebin library for Go. <br>
[![Go Reference](https://pkg.go.dev/badge/github.com/sa111n111/hastebingo.svg)](https://pkg.go.dev/github.com/sa111n111/hastebingo)

There are two ways to read hastebin data

### Sending raw data

```go
package main

import (
	"fmt"

	"github.com/sa111n111/hastebingo"
)

func main() {
	h := hastebingo.Hastebin{}
	h.Post("HI EVERYONE")
	key := h.GetKey()

	fmt.Println("The key: " + key)

	result, err := h.Read(key) // read from key

	if err != nil {
		panic(err)
	}

	fmt.Println(result) 
}

```

### File sending

```go
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
```


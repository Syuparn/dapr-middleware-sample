package main

import (
	"fmt"

	"github.com/wapc/wapc-guest-tinygo"
)

func main() {
	wapc.RegisterFunctions(wapc.Functions{"rewrite": rewrite})
}

// rewrite returns a new URI if necessary.
func rewrite(requestURI []byte) ([]byte, error) {
	fmt.Printf("request: %s", string(requestURI))

	if string(requestURI) == "v1.0/hi" {
		return []byte("v1.0/invoke/hello/method/hello"), nil
	}
	return requestURI, nil
}

package main

import (
	"fmt"

	"github.com/nstoker/stoker.dev/pkg/version"
)

func main() {
	fmt.Printf("Hello, from version %s", version.Version())
}

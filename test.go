package main

import (
	"fmt"
	"github.com/rojack96/gonika/codecs"
)

func main() {

	c := codecs.Codecs{}

	t := c.Decode8([]byte{1, 2, 3, 4})

	fmt.Println(t)

	return
}

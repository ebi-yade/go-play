# go-play

Simple utilities for The Go Playground

## Usage

```go
package main

import (
	"fmt"
	"github.com/ebi-yade/go-play"
)

func main() {
	input := "Hello, playground"
	result, err := doSthWhatYouWant(input)
	play.NoError(err) // fail on error
	play.Debug(input, result) // print detailed values in order

	foo := obtainFoo()
	bar := generateBar()
	baz := calculateBaz(foo, bar)
	play.Debug( // also you can print them in key-value pairs
		play.KV("foo", foo),
		play.KV("bar", bar),
        play.KV("baz", baz),
	)
}

```

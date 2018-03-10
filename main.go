package main

import (
	"fmt"
)

func main() {
	feed, err := fetchBasicFeed()
	fmt.Printf("%+v %+v\n", feed, err)
}

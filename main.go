package main

import (
	"fmt"
)

func main() {
	feed, err := FetchEnhancedFeed()
	fmt.Printf("%+v %+v\n", feed, err)
}

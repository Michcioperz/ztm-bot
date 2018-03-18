package main

import (
	"fmt"
)

func main() {
	events, err := FetchZTMevents()
	fmt.Printf("%+v %+v\n", events, err)
}

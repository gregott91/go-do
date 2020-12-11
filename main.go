package main

import (
	godo "go-do/src"
)

func main() {
	if err := godo.StartApp(); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"plugin"
)

// loads in our test function which imports this class and calls TestPrint()
func main() {
	p, err := plugin.Open("plugins/test-plugin.so")

	if err != nil {
		log.Fatal("Fatal error when opening plugin! Err: ", err)
	} else {
		initFunc, err := p.Lookup("A")

		if err != nil {
			log.Fatal("Error when looking up func! Err: ", err)
		} else {
			initFunc.(func())()
		}
	}
}

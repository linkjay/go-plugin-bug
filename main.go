package main

import (
	"log"
	"plugin"
)

// loads in our test function which imports this class and calls TestPrint()
func main() {
	// open our plugin path
	p, err := plugin.Open("plugins/test-plugin.so")

	// check if we were successfully able to open it
	if err != nil {
		log.Fatal("Fatal error when opening plugin! Err: ", err)
	} else {
		// now that we have it open, let's lookup the symbol A which should represent the func A() found within test-plugin.go
		initFunc, err := p.Lookup("A")

		// see if we found that properly
		if err != nil {
			log.Fatal("Error when looking up func! Err: ", err)
		} else {
			// now this is how we executed that function
			initFunc.(func())()
		}
	}
}

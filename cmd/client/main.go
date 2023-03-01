package main

import (
	"github.com/reiver/go-telnet"
)

func main() {
	var caller telnet.Caller = telnet.StandardCaller

	telnet.DialToAndCall("localhost:8080", caller)
}

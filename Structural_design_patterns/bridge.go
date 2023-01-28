package main

import (
	"fmt"
)

// Bridge is the interface that the client expects to see.
type Bridge interface {
	// Request is the method that the client expects to see.
	Request() string
}

type Implementation interface {
	// SpecificRequest is the method that the client expects to see.
	SpecificRequest() string
}

type ImplementationImpl struct{}

// SpecificRequest is the method that the client expects to see.

func (a *ImplementationImpl) SpecificRequest() string {
	return "Specific Request"
}

type BridgeImpl struct {
	Implementation
}

// Request is the method that the client expects to see.
func (a *BridgeImpl) Request() string {
	return a.SpecificRequest()
}

// main function
func main() {
	bridge := &BridgeImpl{&ImplementationImpl{}}
	fmt.Println(bridge.Request())
}

package main

import (
	"fmt"
)

// Adapter is the interface that the client expects to see.
type Adapter interface {
	// Request is the method that the client expects to see.
	Request() string
}

type Adaptee interface {
	// SpecificRequest is the method that the client expects to see.
	SpecificRequest() string
}

type AdapteeImpl struct{}

// SpecificRequest is the method that the client expects to see.
func (a *AdapteeImpl) SpecificRequest() string {
	return "Specific Request"
}

type AdapterImpl struct {
	Adaptee
}

// Request is the method that the client expects to see.
func (a *AdapterImpl) Request() string {
	return a.SpecificRequest()
}

// main function
func main() {
	adapter := &AdapterImpl{&AdapteeImpl{}}
	fmt.Println(adapter.Request())
}

package mypackage

import "fmt"

type MyType struct {
	EmbeddedType
}
type EmbeddedType string

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}
func (e EmbeddedType) unexportedMethod() {
}

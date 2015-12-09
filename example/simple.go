package example

import (
	"fmt"
	"github.com/mcuadros/go-jsonschema-generator"
)

// ExampleBasic needs to be json schemafied
type ExampleBasic struct {
	// Foo represents a boolean foo
	Foo bool   `json:"foo"`
	Bar string `json:",omitempty"`
	Qux int8
	Baz []string
}

func main() {
	s := &jsonschema.Document{}
	s.Read(&ExampleBasic{})
	fmt.Println(s)
}

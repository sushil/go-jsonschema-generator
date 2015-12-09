go-jsonschema-generator
=======================

Originally based on: https://github.com/mcuadros/go-jsonschema-generator. This version adds support to put "description" field base on the godoc on exported types.

Basic [json-schema](http://json-schema.org/) generator based on Go types, for easy interchange of Go structures across languages.


Installation
------------

The recommended way to install go-jsonschema-generator

```
go get github.com/sushil/go-jsonschema-generator
```

Examples
--------

A basic example:

```go
package example

// ExampleBasic needs to be json schemafied
type ExampleBasic struct {
	// Foo represents a boolean foo
	Foo bool   `json:"foo"`
	Bar string `json:",omitempty"`
	Qux int8
	Baz []string
}
```

and in main.go

```go
func main() {
	s := &jsonschema.Document{}

	// load files to parse godocs and make a map of them
	fileName := "./example/simple.go"
	// load multiple files as ..
	// f2 := "./example/simple2.go"
	// f3 := "./example/simple3.go"

	comment1 := getCommentMap(fileName)
	// comment2 := getCommentMap(f2)
	// comment3 := getCommentMap(f3)

	commentMap := mapUnion(comment1)
	// for multiple files do this ..
	// commentMap := mapUnion(comment1, comment2, comment3)

	jsonschema.SetCommentMap(commentMap)

	i := new(example.ExampleBasic)

	// pass in struct's address to get doc generated
	s.Read(&i)
	fmt.Println(s)

}
```

this prints lines with a list of Types and exported godocs for them followed
by the json for the type specified (see ExampleBasic) above
```json
62,204 --- "ExampleBasic needs to be json schemafied\n"
123,146 --- "Foo represents a boolean foo\n"
{
    "$schema": "http://json-schema.org/schema#",
    "type": "object",
    "properties": {
        "Bar": {
            "type": "string",
            "description": ""
        },
        "Baz": {
            "type": "array",
            "items": {
                "type": "string"
            },
            "description": ""
        },
        "Qux": {
            "type": "integer",
            "description": ""
        },
        "foo": {
            "type": "bool",
            "description": "represents a boolean foo\n"
        }
    },
    "required": [
        "foo",
        "Qux",
        "Baz"
    ],
    "description": ""
}

```

Currently, you'd need to modify main to load comments from additional files if
needed, and change the type you want json schema for.

This mostly works, but sometimes might not be accurate as simple type names
might collide on map, therefore keeping the doc for the last one loaded from
files.

License
-------

MIT, see [LICENSE](LICENSE)

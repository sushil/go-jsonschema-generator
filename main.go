package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/sushil/go-jsonschema-generator/example"
	"github.com/sushil/go-jsonschema-generator/jsonschema"
)

func main() {
	s := &jsonschema.Document{}

	// load files to parse godocs and make a map of them
	fileName := "./example/simple.go"
	// load multiple files as ..
	// f2 := "./example/simple2.go"
	/// f3 := "./example/simple3.go"

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

func mapUnion(maps ...map[string]string) (union map[string]string) {

	union = make(map[string]string)

	for _, m := range maps {
		for k, v := range m {
			union[k] = v
		}
	}

	return union
}

func getCommentMap(fileName string) (cm map[string]string) {
	fileSet := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	f, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	commentMap := ast.NewCommentMap(fileSet, f, f.Comments)
	cm = make(map[string]string)
	for n, cgs := range commentMap {
		fmt.Printf("%#v,%#v --- %#v\n", n.(ast.Node).Pos(), n.(ast.Node).End(), toText(cgs))
		comment := toText(cgs)
		if len(strings.TrimSpace(comment)) == 0 {
			continue
		}
		split := strings.SplitN(comment, " ", 2)
		godoc := split[1]
		key := split[0]
		cm[key] = godoc
	}
	return cm
}

func getComment(c *ast.CommentGroup) string {

	if c == nil {
		return ""
	}

	fmt.Println(c.Text())
	for _, lc := range c.List {
		fmt.Println(lc.Text)
	}

	return c.Text()

}

func toText(cgs []*ast.CommentGroup) (text string) {
	var buf bytes.Buffer
	for _, cg := range cgs {
		buf.WriteString(cg.Text())
	}

	return buf.String()
}

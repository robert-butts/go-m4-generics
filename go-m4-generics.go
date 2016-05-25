//go:generate sh -c "m4 -D TYPE=int $GOPATH/src/github.com/robert-butts/go-m4-generics/binarytree/binarytree.go.m4 > $GOPATH/src/github.com/robert-butts/go-m4-generics/binarytree/intbinarytree.go"
//go:generate sh -c "m4 -D TYPE=string $GOPATH/src/github.com/robert-butts/go-m4-generics/binarytree/binarytree.go.m4 > $GOPATH/src/github.com/robert-butts/go-m4-generics/binarytree/stringbinarytree.go"
package main

import (
	"fmt"
	"github.com/robert-butts/go-m4-generics/binarytree"
)

type CustomType struct {
	I int
	S string
}

func (c CustomType) String() string {
	return fmt.Sprintf(`{"Name":"CustomType", "I":%d, "S", "%s"}`, c.I, c.S)
}

func main() {
	i := binarytree.Newint()
	i.Head.AddLeft(9)
	i.Head.AddRight(19)
	i.Head.Right.AddLeft(994)
	fmt.Println(i)

	s := binarytree.Newstring()
	s.Head.AddLeft("fu")
	s.Head.AddRight("bar")
	s.Head.Right.AddLeft("baz")
	fmt.Println(s)
}
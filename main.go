package main

import "fmt"

type Tree struct {
	length int
	root   *Node
}

type Node struct {
	no int
}

func Info(t Tree) {
	fmt.Printf("tree pointer: %p\n", &t)
	fmt.Printf("pointer: %p\n", t.root)
}

func main() {
	r := &Node{no: 1}
	t := Tree{length: 10, root: r}
	fmt.Printf("tree pointer: %p\n", &t)
	fmt.Printf("root pointer: %p\n", t.root)
	Info(t)
}

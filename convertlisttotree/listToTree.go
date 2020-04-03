package main

import "fmt"

type data struct {
	id     string
	parent string
}

type Node struct {
	Parent *Node
	ID     string
}

func main() {
	testData := []data{
		{
			id:     "bar",
			parent: "baz",
		},
		{
			id:     "foo",
			parent: "",
		},
		{
			id:     "baz",
			parent: "foo",
		},
	}

	m := make(map[string]*Node)
	for _, val := range testData {
		m[val.id] = &Node{ID: val.id}
	}
	for _, val := range testData {
		if val.parent != "" {
			m[val.id].Parent = m[val.parent]
		}
	}
	for _, node := range m {
		if node.Parent != nil {
			fmt.Printf("%v parent is %v\n", node.ID, node.Parent.ID)
		} else {
			fmt.Printf("%v parent is null\n", node.ID)
		}
	}
}

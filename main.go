package main

import "fmt"

type FileNode struct {
	Name     string
	Parent   *FileNode
	Children []*FileNode
	Depth    int
}

func (f *FileNode) AddChild(child *FileNode) *FileNode {
	child.Parent = f
	f.Children = append(f.Children, child)
	child.Depth = f.Depth + 1
	return child
}

func findParent(node1, node2 *FileNode) *FileNode {
	if node1 == nil || node2 == nil {
		return nil
	}

	// use node2 as the deeper node
	if node1.Depth > node2.Depth {
		node1, node2 = node2, node1
	}

	// move node2 up to the same depth as node1
	for node2.Depth > node1.Depth {
		node2 = node2.Parent
	}

	// move both nodes up until they meet
	for node1 != node2 {
		node1 = node1.Parent
		node2 = node2.Parent
	}

	// return the common parent
	return node1
}

func main() {
	root := &FileNode{Name: "root"}
	a := root.AddChild(&FileNode{Name: "a"})
	b := root.AddChild(&FileNode{Name: "b"})

	c := a.AddChild(&FileNode{Name: "c"})
	d := a.AddChild(&FileNode{Name: "d"})

	b.AddChild(&FileNode{Name: "e"})

	fmt.Println(findParent(a, b).Name)
	fmt.Println(findParent(c, d).Name)
}

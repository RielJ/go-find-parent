package main

import "testing"

func TestFindParent(t *testing.T) {
	root := &FileNode{Name: "root"}
	a := root.AddChild(&FileNode{Name: "a"})
	b := root.AddChild(&FileNode{Name: "b"})

	c := a.AddChild(&FileNode{Name: "c"})
	d := a.AddChild(&FileNode{Name: "d"})

	b.AddChild(&FileNode{Name: "e"})

	if findParent(a, b) != root {
		t.Error("Expected root to be the parent of a and b")
	}

	if findParent(c, d) != a {
		t.Error("Expected a to be the parent of c and d")
	}
}

func TestFindParentWithNil(t *testing.T) {
	if findParent(nil, nil) != nil {
		t.Error("Expected nil to be returned when both nodes are nil")
	}

	root := &FileNode{Name: "root"}
	if findParent(root, nil) != nil {
		t.Error("Expected nil to be returned when one node is nil")
	}
}

func TestFindParentWithDifferentDepths(t *testing.T) {
	root := &FileNode{Name: "root"}
	a := root.AddChild(&FileNode{Name: "a"})
	b := root.AddChild(&FileNode{Name: "b"})

	c := a.AddChild(&FileNode{Name: "c"})

	e := b.AddChild(&FileNode{Name: "e"})
	f := e.AddChild(&FileNode{Name: "f"})

	if findParent(c, f) != root {
		t.Error("Expected root to be the parent of c and f")
	}
}

package main

import (
	"testing"
)

func TestSimpleCycleIsDetected(t *testing.T) {
	a := NewNode()
	b := NewNode()
	c := NewNode()
	a.AddChild(b)
	b.AddChild(c)
	c.AddChild(a)

	graph := []*Node{a, b, c}
	if !HasCycle(graph) {
		t.Errorf("Cycle A->B->C->A not detected")
	}
}

func TestNoCycleDetectedInTree(t *testing.T) {
	a := NewNode()
	b := NewNode()
	c := NewNode()
	a.AddChild(b)
	b.AddChild(c)
	a.AddChild(c)

	graph := []*Node{a, b, c}
	if HasCycle(graph) {
		t.Errorf("No cycle present, but dected in A->B, A->C")
	}
}

func TestComplexCyclesWithBranching(t *testing.T) {
	a := NewNode()
	b := NewNode()
	c := NewNode()
	d := NewNode()
	e := NewNode()
	f := NewNode()
	g := NewNode()
	h := NewNode()
	i := NewNode()
	a.AddChild(b)
	b.AddChild(c)
	c.AddChild(d)
	d.AddChild(e)
	e.AddChild(b)
	i.AddChild(h)
	h.AddChild(g)
	f.AddChild(b)
	e.AddChild(h)

	graph := []*Node{a, b, c, d, e, f, g, h, i}
	if !HasCycle(graph) {
		t.Errorf("Complex cycles not detected")
	}
}

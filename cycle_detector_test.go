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

        graph := []*Node{a,b,c}
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

        graph := []*Node{a,b,c}
        if HasCycle(graph) {
            t.Errorf("No cycle present, but dected in A->B, A->C")
        }
}

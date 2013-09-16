package main

import (
    "fmt"
)

type Node struct {
    children []*Node
    index int
    lowlink int
}

func NewNode() *Node {
    node := Node {
        children : []*Node{},
    }
    return &node
}

func (parent *Node) AddChild(child *Node) {
    parent.children = append(parent.children, child)
}
/*
algorithm tarjan is
  input: graph G = (V, E)
  output: set of strongly connected components (sets of vertices)

  index := 0
  S := empty
  for each v in V do
     if (v.index is undefined) then
      strongconnect(v)
    end if
  end for

  function strongconnect(v)
    // Set the depth index for v to the smallest unused index
    v.index := index
    v.lowlink := index
    index := index + 1
    S.push(v)

    // Consider successors of v
     for each (v, w) in E do
       if (w.index is undefined) then
         // Successor w has not yet been visited; recurse on it
        strongconnect(w)
        v.lowlink  := min(v.lowlink, w.lowlink)
      else if (w is in S) then
         // Successor w is in stack S and hence in the current SCC
         v.lowlink  := min(v.lowlink, w.index)
      end if
    end for

    // If v is a root node, pop the stack and generate an SCC
    if (v.lowlink = v.index) then
      start a new strongly connected component
      repeat
        w := S.pop()
        add w to current strongly connected component
      until (w = v)
      output the current strongly connected component
    end if
  end function
*/

func min(a int, b int) (int) {
    if a < b {
        return a
    }
    return b
}

func present(a *Node, list []*Node) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func HasCycle(graph []*Node) (bool) {
        sscs := []([]*Node){}
        s := []*Node{}
        index := 1
        var strongconnect func(v *Node)
        strongconnect = func(v *Node) {
            v.index = index
            v.lowlink = index
            index += 1
            s = append(s, v)
            for _, w := range v.children {
                if w.index == 0 {
                    strongconnect(w)
                    v.lowlink = min(v.lowlink, w.lowlink)
                } else if present(w, s) {
                    v.lowlink = min(v.lowlink, w.index)
                }
            }
            if (v.lowlink == v.index) {
                //Loopy bit
                fmt.Printf("SSC START\n")
                ssc := []*Node{}
                w := s[len(s)-1]
                s = append(s[:len(s)-1])
                fmt.Printf("Adding something to an ssc\n")
                ssc = append(ssc, w)
                for v != w {
                    w = s[len(s)-1]
                    s = append(s[:len(s)-1])
                    fmt.Printf("Adding something to an ssc\n")
                    ssc = append(ssc, w)
                }
                fmt.Printf("SSC END: %v\n", ssc)
                sscs = append(sscs, ssc)
            }
        }
        for _, v := range graph {
            fmt.Printf("%v\n",v.index)
            if v.index == 0 {
                strongconnect(v)
            }
        }
        for _, ssc := range sscs {
            if len(ssc) > 1 {
                return true
            }
        }
    return false;
}

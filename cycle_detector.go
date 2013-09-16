package main

type Node struct {
	children []*Node
	index    int
	lowlink  int
}

func NewNode() *Node {
	node := Node{
		children: []*Node{},
	}
	return &node
}

func (parent *Node) AddChild(child *Node) {
	parent.children = append(parent.children, child)
}

func min(a int, b int) int {
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

func getLargeSSCS(graph []*Node) []([]*Node) {
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
		if v.lowlink == v.index {
			ssc := []*Node{}
			w := s[len(s)-1]
			s = append(s[:len(s)-1])
			ssc = append(ssc, w)
			for v != w {
				w = s[len(s)-1]
				s = append(s[:len(s)-1])
				ssc = append(ssc, w)
			}
			sscs = append(sscs, ssc)
		}
	}
	for _, v := range graph {
		if v.index == 0 {
			strongconnect(v)
		}
	}
	large_sscs := []([]*Node){}
	for _, ssc := range sscs {
		if len(ssc) > 1 {
			large_sscs = append(large_sscs, ssc)
		}
	}
	return large_sscs
}

func HasCycle(graph []*Node) bool {
	large_sscs := getLargeSSCS(graph)
	if len(large_sscs) > 0 {
		return true
	}
	return false
}

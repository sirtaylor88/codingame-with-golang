package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func Find(n *Node, v int) *Node {
	for n != nil {
		if v < n.value {
			n = n.left
		} else if v > n.value {
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

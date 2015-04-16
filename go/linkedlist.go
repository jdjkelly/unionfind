package main

type Node struct {
	data int
	next *Node
}

func main() {
	head := Node{data: 0}
	node := head

	for i, s := range(999) {
		newNode := Node{data: i}
		node.next = newNode
		node = node.next
	}
}
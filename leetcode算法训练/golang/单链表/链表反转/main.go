package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func buildLink(node *Node, val chan int) *Node {
	v, ok := <-val
	if ok == false {
		return nil
	} else {
		node.val = v
		node.next = buildLink(&Node{}, val)
	}
	return node
}

func foreach(node *Node) {

	fmt.Println(node.val)
	if node.next == nil {
		return
	} else {
		foreach(node.next)
	}
	return
}

func reverse(oldNode *Node, newNode *Node, lastNode *Node) (int, *Node) {

	if oldNode.next == nil {
		return oldNode.val, nil
	}
	val, root := reverse(oldNode.next, &Node{}, newNode)
	if root == nil {
		root = &Node{}
		root.next = lastNode
		root.val = val
	} else {
		newNode.next = lastNode
		newNode.val = val
	}
	return oldNode.val, root
}

func main() {
	val := make(chan int, 10)
	for i := 0; i < 10; i++ {
		val <- i
	}
	close(val)

	link := buildLink(&Node{}, val)
	//foreach(link)
	var node = &Node{}
	var root *Node
	node.val, root = reverse(link, &Node{}, node)
	foreach(root)
}

package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{property: property, nextNode: linkedList.headNode}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = Node{property: property, nextNode: nil}

	var lastNode = linkedList.LastNode()
	if lastNode != nil && lastNode.nextNode == nil {
		lastNode.nextNode = &node
	} else {
		fmt.Println("Unable to add node at last")
	}

}

func (linkedList *LinkedList) find(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = Node{property: property, nextNode: nil}
	var prevNode *Node

	prevNode = linkedList.find(nodeProperty)
	if prevNode != nil {
		node.nextNode = prevNode.nextNode
		prevNode.nextNode = &node
	}
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	var linkedList = LinkedList{}
	linkedList.AddToHead(1)
	fmt.Println("Current Head Node : ", linkedList.headNode.property)
	linkedList.AddToHead(2)
	fmt.Println("Current Head Node : ", linkedList.headNode.property)
	linkedList.AddToHead(3)
	fmt.Println("Current Head Node : ", linkedList.headNode.property)
	linkedList.AddToHead(4)
	fmt.Println("Current Head Node : ", linkedList.headNode.property)

	fmt.Println("IterateList :")
	linkedList.IterateList()

	fmt.Println("Last Node : ", linkedList.LastNode().property)

	fmt.Println("IterateList :")
	linkedList.IterateList()

	linkedList.AddToEnd(30)
	fmt.Println("Current Head Node : ", linkedList.headNode.property)
	fmt.Println("Last Node : ", linkedList.LastNode().property)

	fmt.Println("IterateList :")
	linkedList.IterateList()

	fmt.Println("Find :", linkedList.find(4).property)
	fmt.Println("IterateList :")
	linkedList.IterateList()

	linkedList.AddAfter(4, 7)
	fmt.Println("IterateList :")
	linkedList.IterateList()

}

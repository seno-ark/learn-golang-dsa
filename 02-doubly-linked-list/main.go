package main

import "fmt"

type LinkedList struct {
	headNode *Node
}

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// Add a new node to the head of the list
func (l *LinkedList) AddToHead(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}

	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.previousNode = node
	}
	l.headNode = node
}

// getNodeWithValue Get a node with value
func (l *LinkedList) getNodeWithValue(property int) *Node {
	var nodeWith *Node

	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter Add a new node after the following node
func (l *LinkedList) AddAfter(nodeProperty int, property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}

	nodeWith := l.getNodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// getLastNode get last node of the list
func (l *LinkedList) getLastNode() *Node {
	var lastNode *Node

	for node := l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd Add a new node to the end of the list
func (l *LinkedList) AddToEnd(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}

	lastNode := l.getLastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

// iterateList iterates over LinkedList
func iterateList(list LinkedList) {
	for node := list.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.AddToHead(1)   // 1
	linkedList.AddToHead(3)   // 3,1
	linkedList.AddToEnd(5)    // 3,1,5
	linkedList.AddAfter(1, 7) // 3,1,7,5
	linkedList.AddAfter(1, 7) // 3,1,7,7,5
	linkedList.AddAfter(7, 2) // 3,1,7,2,7,5

	iterateList(linkedList)
}

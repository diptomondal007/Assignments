package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	//linkedList.nextNode = nil
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
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

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
		}
	}
	return nodeWith
}

func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}

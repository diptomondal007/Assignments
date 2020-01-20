package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	key       int
	value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinaryTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (tree *BinaryTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	var treeNode *TreeNode
	treeNode = &TreeNode{
		key:       key,
		value:     value,
		LeftNode:  nil,
		RightNode: nil,
	}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {
		insertTreeNode(tree.rootNode, treeNode)
	}
}

func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.LeftNode == nil {
			rootNode.LeftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.LeftNode, newTreeNode)
		}
	} else {
		if rootNode.RightNode == nil {
			rootNode.RightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.RightNode, newTreeNode)
		}
	}
}

func (tree *BinaryTree) SearchElement(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	searchNode(tree.rootNode, key)
}

func searchNode(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return searchNode(node.LeftNode, key)
	}
	if key > node.key {
		return searchNode(node.RightNode, key)
	}
	return true
}

func main() {
	var tree *BinaryTree = &BinaryTree{}
	tree.InsertElement(1, 5)
	tree.InsertElement(2, 4)
	fmt.Println(tree.rootNode.value)
}

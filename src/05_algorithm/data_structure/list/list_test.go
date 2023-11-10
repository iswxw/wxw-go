// @Time : 2023/7/25 15:01
// @Author : xiaoweiwei
// @File : BM0_list_test

package list

import (
	"fmt"
	"testing"
)

// Node 定义节点
type Node struct {
	Val  int   `json:"val"`
	Next *Node `json:"next"`
}

func TestHello(t *testing.T) {

	// 初始化一个链表
	headNode := initList([]int{1, 2, 3, 4, 5})
	fmt.Println("1.初始化一个链表：", showList(headNode))

	// 尾插法新增一个节点
	addLast(headNode, 6)
	fmt.Println("2.尾插法新增一个节点：", showList(headNode))

	// 头插法插入一个节点
	fmt.Println("3.头插法插入一个节点", showList(addFirst(headNode, -1)))
}

// NewNode 初始化一个节点
func newNode(value int) *Node {
	return &Node{Val: value, Next: nil}
}

// initList 初始化一个不带头结点的链表
func initList(values []int) *Node {
	var headNode, curNode *Node
	for _, v := range values {
		node := newNode(v)
		if headNode == nil {
			headNode = node
			curNode = headNode
			continue
		}
		curNode.Next = node
		curNode = curNode.Next
	}
	return headNode
}

// showList 打印链表，不带头结点查询
func showList(node *Node) []int {
	if node == nil {
		return []int{}
	}
	results := make([]int, 0)
	curNode := node
	for curNode != nil {
		results = append(results, curNode.Val)
		curNode = curNode.Next
	}
	return results
}

// addLast 尾插法新增节点
func addLast(headNode *Node, val int) {
	if headNode == nil {
		return
	}
	curNode := headNode
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = newNode(val)
	return
}

// addFirst 头插法，新增节点
func addFirst(headNode *Node, val int) *Node {
	if headNode == nil {
		return headNode
	}
	newNode := newNode(val)
	newNode.Next = headNode
	return newNode
}

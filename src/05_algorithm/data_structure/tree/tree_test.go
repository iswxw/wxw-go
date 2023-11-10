// @Time : 2023/7/27 11:01
// @Author : xiaoweiwei
// @File : tree_test

package tree

// Node 二叉树的节点
type Node struct {
	leftNode  *Node
	rightNode *Node
	val       int
	depth     int
}

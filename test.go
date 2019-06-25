package main

import "fmt"

/**
二叉树：要出输出所有叶子节点，每一个节点的节点号也要输出来
 */
type node struct {
	id int
	left *node
	right *node
}

func getAllMap(node *node) [][][2]int{
	list := make([][][2]int,0)
	var leaf [][2]int = [][2]int{}
	count := 0
	getAllNode(node,leaf,&list,&count)
	return list
}
func getAllNode(node *node, leaf [][2]int, list *[][][2]int,count *int) {
	if node.left == nil && node.right==nil {//边界处理
		*count ++
		leaf = append(leaf,[2]int{*count,node.id})
		*list = append(*list,leaf)
		return
	}
	*count++
	leaf = append(leaf,[2]int{*count,node.id})
	if node.left != nil {//前进条件1
		getAllNode(node.left,leaf,list,count)
	}
	if node.right != nil {//前进条件2
		getAllNode(node.right,leaf,list,count)
	}
}
func main(){
	node1 := node{id:1}
	node2 := node{id:2}
	node3 := node{id:3}
	node4 := node{id:4}
	node5 := node{id:5}
	node6 := node{id:6}
	node7 := node{id:7}
	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	node2.right = &node5
	node3.left = &node6
	node3.right = &node7
//[
// [[1 1] [2 2] [3 4]]
// [[1 1] [2 2] [3 2] [4 5]]
// [[1 1] [2 1] [3 3] [4 3]]
// [[1 1] [2 1] [3 3] [4 3] [5 7]]
// ]
	list := getAllMap(&node1)
	fmt.Println(list)
}

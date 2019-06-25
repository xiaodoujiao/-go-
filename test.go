package main

import "fmt"

/**
二叉树：要求：输出所有叶子节点，每一个节点的节点号也要输出
涉及到的知识点：
1.递归函数调用，把一帧桢数据押入栈，然后最后一个函数调用结束，再一个个桢地释放，押入栈桢的是一些数据
2.这些桢都不是并行的，即使有两个前进段，也是串行的，押入桢的顺序，对这个例子来说，就是先把左节点押入栈
3.边界条件，是没有节点桢可以押入了，就会逐个从后向前弹栈
4.递归函数之间 通过 参数 拷贝传值，把父递归函数的值 传递给子递归函数桢；如果想在不同的桢之间传递数据，除了拷贝，还可以放在堆区，通过指针
5.拷贝只能拷贝给自己的子桢，如果不是自己的子桢，共享数据只能通过指针。
6.一个前进段和2个前进段的区别就是，每个前进段都会分出来自己的子桢，各个前进段之间，就无法通过拷贝传参了，所有逇拷贝传参，只是在一条栈里面说的；
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

	list := getAllMap(&node1)
	fmt.Println(list)
}

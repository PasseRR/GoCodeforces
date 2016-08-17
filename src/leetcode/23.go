package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	temp := lists
	lists = []*ListNode{}
	// 去掉数组中为nil的链表
	for _, item := range temp {
		if item != nil {
			lists = append(lists, item)
		}
	}
	node := &ListNode{}
	current := node
	var length int
	// 移除lists中为nil的链表
	// 每次只会移除一个链表
	var removeNilNode = func() bool {
		for index, item := range lists {
			if item == nil {
				lists = append(lists[:index], lists[index+1:]...)
				return true
			}
		}

		return false
	}

	// 获得数组链表中 头结点最小值
	// 将最小值链表后移动一个节点
	var getMinimumAndMoveNode = func() int {
		min := math.MaxInt32
		minIndex := -1
		for index, item := range lists {
			if item.Val <= min {
				min = item.Val
				minIndex = index
			}
		}

		lists[minIndex] = lists[minIndex].Next

		return min
	}

	for len(lists) > 0 {
		length = len(lists)
		if length > 1 {
			current.Next = &ListNode{}
			current = current.Next
			current.Val = getMinimumAndMoveNode()
			removeNilNode()
		} else {
			// 数组中只有一个非空链表时结束合并
			current.Next = lists[0]
			break
		}
	}

	return node.Next
}

func printNodeArray(lists []*ListNode) {
	for _, item := range lists {
		printNode(item)
	}
}

func printNode(node *ListNode) {
	for i := node; i != nil; i = i.Next {
		fmt.Printf("%v->", i.Val)
	}
	fmt.Printf("\n")
}

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
		},
	}

	node2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 55,
			},
		},
	}

	lists := []*ListNode{node1, node2}

	printNode(mergeKLists(lists))
}

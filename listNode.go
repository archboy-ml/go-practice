package main

import "fmt"

type Object interface {}

type Node struct {
	Data Object // 定义数据域
	Next *Node // 定义指针域
}

type List struct {
	headNode *Node // 头节点
}

func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	} else {
		return false
	}
}

func (l *List) Length() int {
	// 获取链表头节点
	cur := l.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (l *List) Shift(data Object) *Node {
	node := &Node{Data:data}
	node.Next = l.headNode
	l.headNode = node
	return node
}

func (l *List) Append(data Object) {
	node := &Node{Data:data}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		cur := l.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (l *List) ShowList() {
	if !l.IsEmpty() {
		cur := l.headNode
		for {
			fmt.Printf("%v\n", cur.Data)
			if  cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	} else {
		fmt.Println("List为空")
	}
}

func reverseList(list List) {

}

func main() {
	var list List
	for i:=1;i<6;i++ {
		list.Append(i)
	}
	list.ShowList()
}

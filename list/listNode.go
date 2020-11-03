package linkedList

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

func (l *List) Insert(index int, data Object)  {
	if index <= 0 {
		l.Shift(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		pre := l.headNode
		count := 0
		for count < (index-1) {
			pre = pre.Next
			count++
		}
		node := &Node{Data:data}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (l *List) Remove(data Object) {
	pre := l.headNode
	if pre.Data == data {
		l.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

func (l *List) RemoveAtIndex(index int) {
	pre := l.headNode
	if index <= 0 {
		l.headNode = pre.Next
	} else if index > l.Length() {
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0
		for count < (index-1) && pre.Next != nil {
			pre = pre.Next
			count++
		}
		pre.Next = pre.Next.Next
	}
}

func (l *List) Contain(data Object) bool {
	cur := l.headNode
	for cur.Next != nil {
		if cur.Data == data {
			return true
		} else {
			cur = cur.Next
		}
	}
	return false
}

func (l *List) ShowList() {
	if !l.IsEmpty() {
		cur := l.headNode
		for {
			fmt.Printf("%v\t", cur.Data)
			if  cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
		fmt.Println()
	} else {
		fmt.Println("List为空")
	}
}

func ReverseList(list List) List {
	head := list.headNode
	var result List
	if head == nil {
		fmt.Println("List为空")
		result = List{}
	} else {
		node := &Node{Data:head.Data}
		cur := head.Next
		for cur != nil {
			current := cur
			cur = cur.Next
			current.Next = node
			node = current
		}
		result.headNode = node
	}
	//fmt.Printf("%v\n", result)
	return result
}

func main() {
	var list List
	for i:=1;i<6;i++ {
		list.Append(i)
	}
	list.ShowList()
	println("================================")
	var l List
	l = ReverseList(list)
	l.ShowList()
}

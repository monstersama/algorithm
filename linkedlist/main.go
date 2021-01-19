package main

import "fmt"

type Item interface{}

//节点结构体
type ListNode struct {
	Data Item
	Next *ListNode
}

// 链表操作接口
type LinkedList interface {
	Add(data Item)  // 头插法
	Append(data Item)  // 尾插法
	Insert(index uint, data Item)
	Delete(index uint) Item
	Len() uint
	GetIndex(data Item) uint
	GetNode(index uint) Item
	GetAll()
	//Reverse() *ListNode
}

// 尾插
func (l *ListNode) Append(data Item) {
	node := l
	for node.Next != nil {
		node = node.Next
	}
	newNode := ListNode{data, nil}
	node.Next = &newNode
}

// 头插
func (l *ListNode) Add(data Item) {
	node := l
	newNode := ListNode{data, nil}
	newNode.Next = node.Next
	node.Next = &newNode
}

//获取链表长度
func (l *ListNode) Len() uint {
	node := l
	var length uint
	for node.Next != nil {
		node = node.Next
		length++
	}
	return length
}

//任意位置插入
func (l *ListNode) Insert(index uint, data Item) bool {
	node := l
	length := node.Len()
	if length < index {
		return false
	}
	for i := 0; uint(i) < index; i++ {
		node = node.Next
	}
	newNode := ListNode{Data: data}
	newNode.Next = node.Next
	node.Next = &newNode
	return true
}

//删除任意位置节点
func (l *ListNode) Delete(index uint) bool {
	node := l
	length := node.Len()
	if length < index {
		return false
	}
	for i := 0; uint(i) < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next
	return true
}

// 遍历节点
func (l *ListNode) GetAll() {
	node := l
	for node.Data != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
}

//根据位置找到节点
func (l *ListNode) GetIndex(data Item) int {
	node := l
	var index int
	for node.Next != nil {
		if node.Data == data {
			return index
		}
		node = node.Next
		index++
	}
	if node.Data == data {
		return index
	}
	return -1
}

func (l *ListNode) GetNode(index uint) Item {
	node := l
	length := node.Len()
	if index > length {
		return nil
	}
	for i := 0; i < int(index); i++ {
		node = node.Next
	}
	target := node.Next.Data
	return target
}

//创建一个链表
func NewLinkedList(length int) *ListNode {
	if length < 0 {
		fmt.Println("Linked list Length must be greater than or equal to 0")
		return nil
	}
	var head *ListNode

	head = &ListNode{}
	for i := 0; i<length; i++ {
		var newNode *ListNode
		newNode = &ListNode{Data: i}
		newNode.Next = head
		head = newNode
	}
	return head
}

func main() {
	length := 10
	node := NewLinkedList(length)
	node.GetAll()
}
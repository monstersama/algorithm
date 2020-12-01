package main

type Object interface{}

type Node struct {
	data Object
	next *Node
}

type List struct {
	size uint64 // 链表长度
	head *Node  // 头指针
	tail *Node  // 尾指针
}

// 初始化链表方法
func (list *List) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).head = nil
}

//添加结点方法
func (list *List) Append(node *Node) {
	if (*list).size == 0 {
		(*list).head = node
		(*list).tail = node
		(*list).size = 1
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node //旧的尾部结点的后继指针指向新的结点
		(*list).tail = node    //链表的尾指针指向新的结点
		(*list).size++         //链表长度自增
	}
}

/*上面的链表结点添加的方法不完善，无法处理空节点，代码有重复*/
func (list *List) Append2(node *Node) bool {
	if node == nil {
		return false
	}

	(*node).next = nil
	if (*list).size == 0 {
		(*list).head = node
	} else {
		oldTail := (*list).tail
		(*oldTail).next = node
		(*list).size++
	}
	return true
}

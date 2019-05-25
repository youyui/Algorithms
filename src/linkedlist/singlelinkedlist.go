package linkedlist

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

type SingleLinekedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{v, nil}
}

func (this *SingleLinekedList) addNode(v interface{}) bool {
	return false
}

package stack

type ListNode struct {
	Next *ListNode
	Val  interface{}
}

type LinkedListStack struct {
	topNode *ListNode
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) IsEmpty() bool {
	if this.topNode == nil {
		return true
	}
	return false
}

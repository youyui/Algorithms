package queue

//循环队列
type CycleQueue struct {
	data     []interface{}
	capacity uint
	head     uint
	tail     uint
}

//创建循环队列！
func NewCycleQueue(cap uint) *CycleQueue {

	return &CycleQueue{
		data:     nil,
		capacity: cap,
		head:     0,
		tail:     0,
	}
}

//入队
func (this *CycleQueue) Push(v interface{}) bool {
	//队列满了
	if (this.tail+1)%this.capacity == this.head {
		return false
	}
	this.data[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity
	return true
}

//出队
func (this *CycleQueue) Poll() interface{} {
	//空队列返回false
	if this.head == this.tail {
		return false
	}
	v := this.data[this.head]

	this.head = (this.head + 1) % this.capacity
	return v
}

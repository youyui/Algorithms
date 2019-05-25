package stack

type ArrayStack struct {
	//数据
	data []interface{} //可以利用slice自动扩容
	//栈顶指针
	top int
}

//新建一个栈！
func NewArrayStack(capacity uint) *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, capacity), //初始化容量
		top:  -1,
	}
}

//判断是否为空
func (this *ArrayStack) IsEmpty() bool {
	if this.top < 0 {
		return true
	}

	return false
}

func (this *ArrayStack) Push(v interface{}) {
	this.top++
	if this.top >= len(this.data) {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}

	v := this.data[this.top]
	this.top--
	return v
}

func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}

	return this.data[this.top]
}

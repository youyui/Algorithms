package array

import (
	"errors"
)

/*
* 实现自动扩容
* 实现大小固定的有序数组，支持动态增删改
*
 */
type Array struct {
	data   []int
	length uint
}

//为数组进行初始化
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

func (this *Array) Len() uint {
	return this.length
}

func (this *Array) isIndexOutOfRange(idx uint) bool {
	if idx >= uint(cap(this.data)) {
		return true
	}
	return false
}

func (this *Array) Find(idx uint) (int, error) {
	if this.isIndexOutOfRange(idx) {
		return 0, errors.New("out of range")
	}
	return this.data[idx], nil
}

func (this *Array) InsertToIdx(idx uint, v int) error {
	if this.Len() == uint(cap(this.data)) {
		return errors.New("full")
	}
	if idx != this.length && this.isIndexOutOfRange(idx) {
		return errors.New("out of range")
	}

	for i := this.length; i > idx; i-- {
		this.data[i] = this.data[i-1]
	}

	this.data[idx] = v
	this.length++

	return nil
}

// 剑指 Offer 09. 用两个栈实现队列（https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/）

package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	obj.AppendTail(2)
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

type CQueue struct {
	inStack, outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		this.in2out()
	}
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}

func (this *CQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

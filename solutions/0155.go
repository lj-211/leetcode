package solutions

import (
//"reflect"
)

type MinStack struct {
	Stack  []int
	Assist []int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	size := len(this.Assist)
	if size > 0 {
		curMin := this.Assist[size-1]
		if x <= curMin {
			this.Assist = append(this.Assist, x)
		}
	} else {
		this.Assist = append(this.Assist, x)
	}
}

func (this *MinStack) Pop() {
	size := len(this.Stack)
	if size == 0 {
		return
	}

	top := this.Stack[size-1]
	this.Stack = this.Stack[0 : size-1]

	// check & pop min
	asize := len(this.Assist)
	if asize > 0 {
		minTop := this.Assist[asize-1]
		if minTop == top {
			this.Assist = this.Assist[0 : asize-1]
		}
	}
}

func (this *MinStack) Top() int {
	size := len(this.Stack)
	if size > 0 {
		return this.Stack[size-1]
	}

	return -1
}

func (this *MinStack) GetMin() int {
	size := len(this.Assist)
	if size > 0 {
		return this.Assist[size-1]
	}

	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func init() {
	desc := `
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
	`
	sol := Solution{
		Title: "最小栈",
		Desc:  desc,
		//Method: reflect.ValueOf(twoSum),
		Tests: make([]TestCase, 0),
	}

	SolutionMap["0155"] = sol
}

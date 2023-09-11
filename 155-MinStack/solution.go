package main

type MinStack struct {
    top *StackNode
    min int
}

type StackNode struct {
    data int
    next *StackNode
    lastmin int
}

var myStack MinStack = MinStack{top: nil}
var newTop *StackNode

func Constructor() MinStack {
    return myStack
}

func (this *MinStack) Push(val int)  {
    if this.top == nil {
        newTop = &StackNode{data: val, next: this.top}
        this.min = val
    } else {
        newTop = &StackNode{data: val, next: this.top, lastmin: this.min}
    }
    this.top = newTop
    if this.top.data < this.min {
        this.min = this.top.data 
    }
}

func (this *MinStack) Pop()  {
    if this.top.next == nil {
        this.top = nil
        return
    }    
    this.min = this.top.lastmin
    *this.top = *this.top.next
}

func (this *MinStack) Top() int {
    return this.top.data    
}

func (this *MinStack) GetMin() int {
    return this.min    
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

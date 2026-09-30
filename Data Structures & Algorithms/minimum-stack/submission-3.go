type MinStack struct {
	items    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.items) == 0 {
		this.minStack = append(this.minStack, val)
	} else if this.minStack[len(this.minStack)-1] >= val {
		this.minStack = append(this.minStack, val)
	}

	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}
	if this.items[len(this.items)-1] == this.minStack[len(this.minStack)-1]{
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.items = this.items[:len(this.items)-1]
	return
}

func (this *MinStack) Top() int {
	var zero int
	if len(this.items) == 0 {
		return zero
	}
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	var zero int
	if len(this.items) == 0 {
		return zero
	}
	return this.minStack[len(this.minStack)-1]
}
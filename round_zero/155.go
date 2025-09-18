package roundzero

type MinStack struct {
	l        int
	stack    []int
	minStack []int
}

func Constructor155() MinStack {
	return MinStack{
		l:        0,
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {

	this.stack = append(this.stack, val)
	if this.l == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(val, this.GetMin()))
	}

	this.l++
}

func (this *MinStack) Pop() {
	if this.l == 0 {
		return
	}

	defer func() {
		this.l--
	}()

	this.stack = this.stack[:this.l-1]
	this.minStack = this.minStack[:this.l-1]
}

func (this *MinStack) Top() int {

	if this.l == 0 {
		return -1
	}

	return this.stack[this.l-1]

}

func (this *MinStack) GetMin() int {
	return this.minStack[this.l-1]
}

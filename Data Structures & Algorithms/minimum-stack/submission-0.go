type MinStack struct {
	data []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)

	if len(this.mins) == 0 || val <= this.mins[len(this.mins) - 1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	lastIdx := len(this.data) - 1
	this.data = this.data[:lastIdx]

	minIdx := len(this.mins) - 1
	if top == this.minTop() {
		this.mins = this.mins[:minIdx]
	}
}

func (this *MinStack) Top() int {
	lastIndex := len(this.data) - 1
	return this.data[lastIndex]
}

func (this *MinStack) minTop() int {
	lastIndex := len(this.mins) - 1
	return this.mins[lastIndex]
}

func (this *MinStack) GetMin() int {
	return this.minTop()
}

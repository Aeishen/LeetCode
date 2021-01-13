package queue

//队列的最大值

type MaxQueue struct {
	normalQueue []int  // 普通队列
	maxValQueue []int  // 单调递减队列
}


func Constructor() *MaxQueue {
	return &MaxQueue{
		make([]int,0),
		make([]int,0),
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.maxValQueue) > 0 {
		return this.maxValQueue[0]
	}
	return -1
}


func (this *MaxQueue) Push_back(value int)  {
	this.normalQueue = append(this.normalQueue, value)
	n :=  len(this.maxValQueue)
	if n > 0 {
		cutIndex := -1
		for i := 0; i < n; i ++ {
			if this.maxValQueue[i] < value {
				cutIndex = i
				break
			}
		}
		if cutIndex != -1 {
			this.maxValQueue[cutIndex] = value
			this.maxValQueue = this.maxValQueue[:cutIndex + 1]
			return
		}
	}
	this.maxValQueue = append(this.maxValQueue, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.normalQueue) > 0 && len(this.maxValQueue) > 0 {
		target := this.normalQueue[0]
		this.normalQueue = this.normalQueue[1:]
		if target == this.maxValQueue[0]{
			this.maxValQueue = this.maxValQueue[1:]
		}
		return target
	}
	return -1
}

/*

*/
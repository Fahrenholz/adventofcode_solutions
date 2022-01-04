package main

type queueItem struct {
	pos       point
	riskLevel int
	idx       int //needed for heap
}

type PriorityQueue []queueItem

// implement heap interface

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].riskLevel < q[j].riskLevel
}

func (q PriorityQueue) Swap(i, j int) {
	q[j], q[i] = q[i], q[j]
	q[i].idx = i
	q[j].idx = j
}

func (q *PriorityQueue) Push(x interface{}) {
	item := x.(queueItem)
	item.idx = len(*q)
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() interface{} {
	before := *q
	res := before[len(*q)-1]
	*q = before[0 : len(*q)-1]
	return res
}

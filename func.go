package main

type Queue []interface{}

//IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//Enqueue
func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

//Dequeue
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

// 제곱함수
func power(a int, b int) int {
	c := 1
	for i := 0; i <= b; i++ {
		c *= a
	}
	return c
}

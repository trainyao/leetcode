package lib

type Queue []interface{}

func (q *Queue) enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) equeue()(v interface{}) {
	v = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) len () int {
	return len(*q)
}

func (q *Queue) empty() bool {
	return len(*q) == 0
}

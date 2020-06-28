package queue

type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Equeue() (v interface{}) {
	if len(*q) == 0 {
		return nil
	}

	v = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) len() int {
	return len(*q)
}

func (q *Queue) empty() bool {
	return len(*q) == 0
}

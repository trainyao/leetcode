package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queue(t *testing.T) {
	q := Queue{}

	q.enqueue(1)
	if !assert.Falsef(t, q.empty(), "queue is not empty") {
		t.FailNow()
	}
	if !assert.Equalf(t, 1, q.len(), "queue len is 1") {
		t.FailNow()
	}

	q.enqueue(2)
	if !assert.Equalf(t, 2, q.len(), "queue len is 1") {
		t.FailNow()
	}

	v := q.equeue()
	if !assert.Equalf(t, 1, v, "value should be 1") {
		t.FailNow()
	}
	if !assert.Equalf(t, 1, q.len(), "queue len is 1") {
		t.FailNow()
	}

	v = q.equeue()
	if !assert.Equalf(t, 2, v, "value should be 2") {
		t.FailNow()
	}
	if !assert.Truef(t, q.empty(), "queue should be empty") {
		t.FailNow()
	}
}

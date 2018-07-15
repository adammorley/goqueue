package queue

import "testing"

func TestQueue(t *testing.T) {
    var q *Queue = New()
    if q.Size() != 0 {
        t.Error("new queue is not right sized")
    }
    q.Enqueue(1)
    if q.Size() != 1 {
        t.Error("wrong size")
    }
    q.Enqueue(1)
    if q.Size() != 2 {
        t.Error("wrong size")
    }
    q.Enqueue(2)
    if q.Size() != 3 {
        t.Error("wrong size")
    }
    if v, e := q.Dequeue(); e != nil || v != 1 {
        t.Error("dequeue not working")
    } else if v, e := q.Dequeue(); e != nil || v != 1 {
        t.Error("dequeue not working")
    } else if v, e := q.Dequeue(); e != nil || v != 2 {
        t.Error("dequeue not working")
    } else if v, e := q.Dequeue(); e == nil || v != 0 {
        t.Error("dequeue not detecting empty queue")
    }
}

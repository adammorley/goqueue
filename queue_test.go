package queue

import "testing"

func TestQueue(t *testing.T) {
    var q *Queue = New(128)
    if q.Size() != 0 {
        t.Error("new queue is not right sized")
    }
    if q.Enqueue(1) != nil {
        t.Error("error on enqueue when not expected")
    } else if q.Size() != 1 {
        t.Error("wrong size")
    }
    if q.Enqueue(1) != nil {
        t.Error("error on enqueue when not expected")
    } else if q.Size() != 2 {
        t.Error("wrong size")
    }
    if v, e := q.Dequeue(); e != nil || v != 1 {
        t.Error("dequeue not working")
    } else if v, e := q.Dequeue(); e != nil || v != 1 {
        t.Error("dequeue not working")
    } else if v, e := q.Dequeue(); e == nil || v != nil {
        t.Error("dequeue not detecting empty queue")
    }

    j := 3
    var q1 *Queue = New(uint(j))
    for i := 0; i < j; i++ {
        q1.Enqueue(1)
    }
    if q1.Enqueue(1) == nil {
        t.Error("expected error about full queue")
    }

    var q2 *Queue = New(4)
    q2.Enqueue(2)
    q2.Enqueue(-2)
    e := q2.Enqueue(4.3)
    // this means: if the error returned is not of the error type defined in the file, then it's an error.  golang syntax, ftw
    if _, ok := e.(*EnqueueTypeError); !ok {
        t.Error("expected error")
    }
}

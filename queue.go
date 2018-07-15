// a queue implementation for integers in golang
package queue

import "errors"

// a queue, which has a pointer to the front and back.  elements are added ("enqueue"'d) to the back and removed ("dequeue"'d) from the front
type Queue struct {
    front *element
    back *element
}

// an element on the queue; each element has a pointer to the next and previous elements in the queue, along with the value stored at the element
type element struct {
    next *element
    prev *element
    value int
}

// a new queue, front and back pointers are nil
func New() *Queue {
    var q *Queue = new(Queue)
    q.front = nil
    q.back = nil
    return q
}

// enqueue a value on the queue
func (q *Queue) Enqueue(v int) {
    var e *element = new(element)
    e.next = nil
    e.value = v
    e.prev = q.back
    if q.front == nil {
        q.front = e
        q.back = e
    } else {
        q.back.next = e
        q.back = e
    }
}

// dequeue a value on the queue.  returns error if queue empty
func (q *Queue) Dequeue() (int, error) {
    if q.front == nil {
        return 0, errors.New("empty queue")
    }
    var e *element = q.front
    q.front = e.next
    return e.value, nil
}

// size of the queue, in number of elements
func (q *Queue) Size() (c uint) {
    var e *element = q.front
    for e != nil {
        c++
        e = e.next
    }
    return
}

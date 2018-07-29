// a queue implementation for values (using interface{} type to allow strings or integers or whatever) in golang.  allows for fixed size queues, and keeps track of the count.  this can be used as a ring buffer
// note that the Enqueue method enforces that each value added is of the same type, as a safety measure
package queue

import "fmt"
import "reflect"

// a queue, which has a pointer to the front and back.  elements are added ("enqueue"'d) to the back and removed ("dequeue"'d) from the front
type Queue struct {
    front *element
    back *element
    size uint
    capacity uint
}

// a queue can return an error message in exceptional conditions
type QueueError struct {
    msg string
}
func (e *QueueError) Error() string {
    return fmt.Sprintf("queue %v", e.msg)
}

// an element on the queue; each element has a pointer to the next and previous elements in the queue, along with the value stored at the element
type element struct {
    next *element
    prev *element
    value interface{}
}

// a new queue, front and back pointers are nil
func New(capacity uint) *Queue {
    var q *Queue = new(Queue)
    q.front = nil
    q.back = nil
    q.capacity = capacity
    q.size = 0
    return q
}

// since a queue can contain any arbitrary type, it can return an error if the type is not the same.
type EnqueueTypeError struct {
    passed string
    stored string
}
func (e *EnqueueTypeError) Error() string {
    return fmt.Sprintf("type mismatch, received: %v, expected: %v", e.passed, e.stored)
}

// enqueue a value on the queue, ensuring that the type of the to-be-added element is the same as the existing elements on the queue.
func (q *Queue) Enqueue(i interface{}) error {
    if q.size + 1 > q.capacity {
        return &QueueError{msg: "full"}
    }
    var e *element = new(element)
    e.next = nil
    e.value = i
    e.prev = q.back
    if q.front == nil {
        q.front = e
        q.back = e
    } else {
        if reflect.TypeOf(i) != reflect.TypeOf(q.front.value) {
            return &EnqueueTypeError{passed: reflect.TypeOf(i).Name(), stored: reflect.TypeOf(q.front.value).Name()}
        }
        q.back.next = e
        q.back = e
    }
    q.size += 1
    return nil
}

// dequeue a value on the queue.  returns error if queue empty
func (q *Queue) Dequeue() (interface{}, error) {
    if q.front == nil {
        return nil, &QueueError{msg: "empty"}
    }
    var e *element = q.front
    q.front = e.next
    q.size -= 1
    return e.value, nil
}

// size of the queue, in number of elements
func (q *Queue) Size() uint {
    return q.size
}

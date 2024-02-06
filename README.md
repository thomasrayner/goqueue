# goqueue

A simple implementation of queues and stacks using slices

## Usage

```go
import (
    "github.com/thomasrayner/goqueue"
)

func main() {
    // Create a new queue
    q := goqueue.NewQueue()
    q.Enqueue("hello")

    // Create a new stack
    s := goqueue.NewStack()
    s.Push("world")
}
```

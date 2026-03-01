# Data Structures

Common data structures implemented from scratch in Go using generics.

## Implementations

### Stack
Generic LIFO (Last In, First Out) stack backed by a slice.
- `Push(elem T)` - add an element to the top
- `Pop() (T, bool)` - remove and return the top element

### Queue
Generic FIFO (First In, First Out) queue backed by a slice.
- `Queue(elem T)` - enqueue an element
- `Dequeue() (T, bool)` - dequeue and return the front element

### HashMap
Hash map with string keys, chaining for collision resolution, and automatic resizing.
- `Set(key string, value any)` - insert or update a key-value pair
- `Get(key string) (any, bool)` - retrieve a value by key
- `Delete(key string) bool` - remove a key-value pair
- `Print()` - display the internal bucket structure

## Usage

```go
s := stack.Build[int](0, 5)
s.Push(10)
s.Push(20)
val, ok := s.Pop() // 20, true

q := queue.Build[string](0, 5)
q.Queue("hello")
q.Queue("world")
val, ok := q.Dequeue() // "hello", true

m := mapa.Build(16)
m.Set("name", "Joel")
v, ok := m.Get("name") // "Joel", true
```

## Run

```bash
go run .
```

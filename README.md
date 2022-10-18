# go_datastructures

## Download
```plaintext
go get github.com/suhashm/go_datastructures
```

## Usage

```plaintext
import (
    stack "github.com/suhashm/go_datastructures/stack"
    queue "github.com/suhashm/go_datastructures/queue"
)

st := stack.New[int](0)
st.Push(1)

q := queue.New[string]("==)
q.Add("A")
```
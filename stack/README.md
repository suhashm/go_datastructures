## Usage

```plaintext
go get github.com/suhashm/go_datastructures
```

```plaintext
import (
    stack "github.com/suhashm/go_datastructures/stack"
)

st := stack.New[int](0)

OR

st := stack.NewWithCapacity[int](5, 0)

st.Push("A")
```
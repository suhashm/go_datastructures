package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOperations(t *testing.T) {
	defaultValue := -1
	stack := New[int](defaultValue)
	stack.Push(1)
	stack.Push(2)
	value, err := stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)

	value, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 2, value)

	value, err = stack.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, value)

	// Pop on an empty stack throws an error
	_, _ = stack.Pop()
	value, err = stack.Pop()
	assert.NotNil(t, err)
	assert.Equal(t, defaultValue, value)
}

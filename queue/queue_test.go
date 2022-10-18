package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueOperation(t *testing.T) {
	defaultValue := "=="
	queue := New[string](defaultValue)

	queue.Add("m")
	queue.Add("n")
	value, err := queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "m", value)

	value, err = queue.Remove()
	assert.Nil(t, err)
	assert.Equal(t, "m", value)

	queue.Add("o")
	value, err = queue.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "n", value)

	// Remove on empty queue throws an error
	_, _ = queue.Remove()
	_, _ = queue.Remove()
	
	res, err := queue.Remove()
	assert.NotNil(t, err)
	assert.Equal(t, defaultValue, res)
	

	
}
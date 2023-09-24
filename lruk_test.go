package lruk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUK(t *testing.T) {
  l := New[int, int](2, 2)

  l.Put(1, 1)
  v, ok := l.Get(1)
  assert.False(t, ok)
  assert.Equal(t, 0, v)

  l.Put(1, 1)
  v, ok = l.Get(1)
  assert.True(t, ok)
  assert.Equal(t, 1, v)

  l.Put(2, 2)
  v, ok = l.Get(2)
  assert.False(t, ok)
  assert.Equal(t, 0, v)

  l.Put(2, 2)
  v, ok = l.Get(2)
  assert.True(t, ok)
  assert.Equal(t, 2, v)

  l.Put(3, 3)
  v, ok = l.Get(3)
  assert.False(t, ok)
  assert.Equal(t, 0, v)

  l.Put(3, 3)
  v, ok = l.Get(3)
  assert.True(t, ok)
  assert.Equal(t, 3, v)

  v, ok = l.Get(1)
  assert.False(t, ok)
  assert.Equal(t, 0, v)
}

package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_PushFront(t *testing.T) {
	l := New[int]()
	tests := []struct {
		input int
		want  int
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
	}

	for _, test := range tests {
	  got := l.PushFront(test.input)
	  assert.Equal(t, test.want, got.Value)
	}

	assert.Equal(t, 1, l.Back().Value)
}

func TestList_Remove(t *testing.T) {
  l := New[string]()
  l.PushFront("1")
  ele := l.PushFront("2")
  l.PushFront("3")

  assert.Equal(t, 3, l.Len())

  l.Remove(ele)
  assert.Equal(t, 2, l.Len())
}

func TestList_MoveToFront(t *testing.T) {
  l := New[string]()
  l.PushFront("1")
  ele := l.PushFront("2")
  l.PushFront("3")

  l.MoveToFront(ele)
  assert.Equal(t, ele, l.root.next)
}

package list

type Element[V any] struct {
	next, prev *Element[V]
	list       *List[V]
	Value      V
}

type List[V any] struct {
	root Element[V]
	len  int
}

func (l *List[V]) Init() *List[V] {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New[V any]() *List[V] {
	return new(List[V]).Init()
}

func (l *List[V]) Len() int {
	return l.len
}

// lazyInit lazily initializes a zero List value.
func (l *List[V]) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List[V]) remove(e *Element[V]) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}

func (l *List[V]) Remove(e *Element[V]) V {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// insert inserts e after at, increments l.len, and returns e.
func (l *List[V]) insert(e, at *Element[V]) *Element[V] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

func (l *List[V]) insertValue(v V, at *Element[V]) *Element[V] {
	return l.insert(&Element[V]{Value: v}, at)
}

func (l *List[V]) PushFront(v V) *Element[V] {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// move moves e to the next to at.
func (l *List[V]) move(e, at *Element[V]) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

func (l *List[V]) MoveToFront(e *Element[V]) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *List[V]) Back() *Element[V] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

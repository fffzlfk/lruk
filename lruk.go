package lruk

import "lruk/list"

type LRUKCache[K comparable, V any] struct {
	k          int
	capacity   int
	cacheMap   map[K]*list.Element[entry[K, V]]
	cache      *list.List[entry[K, V]]
	historyMap map[K]*list.Element[histroyCounter[K]]
	history    *list.List[histroyCounter[K]]
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

type histroyCounter[K comparable] struct {
	key     K
	visited int
}

func New[K comparable, V any](k int, capacity int) *LRUKCache[K, V] {
	return &LRUKCache[K, V]{
		k:          k,
		capacity:   capacity,
		cacheMap:   make(map[K]*list.Element[entry[K, V]], capacity),
		cache:      list.New[entry[K, V]](),
		historyMap: make(map[K]*list.Element[histroyCounter[K]]),
		history:    list.New[histroyCounter[K]](),
	}
}

func (l *LRUKCache[K, V]) Get(key K) (value V, ok bool) {
	ele, ok := l.cacheMap[key]
	if ok {
		value = ele.Value.value
		l.cache.MoveToFront(ele)
		return
	}
	
	hcEle, ok := l.historyMap[key]
	if ok {
	  hcEle.Value.visited++
	  return value, false
	}
	return
}

func (l *LRUKCache[K, V]) addToCache(key K, value V) {
	ele := l.cache.PushFront(entry[K, V]{
		key:   key,
		value: value,
	})
	l.cacheMap[key] = ele
	if l.cache.Len() > l.capacity {
		backEle := l.cache.Back()
		l.cache.Remove(backEle)
		delete(l.cacheMap, backEle.Value.key)
	}
}

func (l *LRUKCache[K, V]) addToHistory(key K, value V) {
	hcEle, ok := l.historyMap[key]
	if ok {
		hc := hcEle.Value
		hc.visited++
		if hc.visited >= l.k {
			l.history.Remove(hcEle)
			delete(l.historyMap, hc.key)
			l.addToCache(key, value)
		}
		return
	}
	ele := l.history.PushFront(histroyCounter[K]{
	  key: key,
	  visited: 1,
	})
  l.historyMap[key] = ele
  if l.history.Len() > 2 * l.capacity {
    back := l.history.Back()
    l.history.Remove(back)
    delete(l.historyMap, back.Value.key)
  }
}

func (l *LRUKCache[K, V]) Put(key K, value V) {
	ele, ok := l.cacheMap[key]
	if ok {
		ele.Value.value = value
		l.cache.MoveToFront(ele)
		return
	}
	
	l.addToHistory(key, value)
}

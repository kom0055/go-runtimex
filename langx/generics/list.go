package generics

import "sync"

type ListWrap[T any] interface {
	Append(t ...T)
	FetchAll() []T
}

func NewListWrap[T any]() ListWrap[T] {
	return &listWrap[T]{}
}

type listWrap[T any] struct {
	items []T
	sync.RWMutex
}

func (l *listWrap[T]) Append(t ...T) {
	l.Lock()
	defer l.Unlock()
	l.items = append(l.items, t...)
}

func (l *listWrap[T]) FetchAll() []T {
	l.RLock()
	defer l.RUnlock()
	return l.items
}

package main

type GenericList[T comparable] struct {
	Data []T
}

func NewGenericList[T comparable]() *GenericList[T] {
	return &GenericList[T]{Data: []T{}}
}

func (l *GenericList[T]) Insert(value T) {
	l.Data = append(l.Data, value)
}

func (l *GenericList[T]) ValueByIndex(index int) T {
	if index >= len(l.Data) {
		panic("Index out of Range")
	}
	return l.Data[index]
}

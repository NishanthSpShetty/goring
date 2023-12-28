package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("empty buffer")
	ErrFull  = errors.New("buffer full")
)

type RingBuf[T any] struct {
	storage []T

	readIdx  uint64
	writeIdx uint64
	full     bool
}

// New Create a new instance of ring buffer of type T
func New[T any](cap int) (*RingBuf[T], error) {
	if cap <= 0 {
		return nil, errors.New("capacity must be non negative")
	}
	return &RingBuf[T]{
		storage:  make([]T, cap),
		readIdx:  0,
		writeIdx: 0,
		full:     false,
	}, nil
}

func (r *RingBuf[T]) Offer(item T) error {
	if r.isFull() {
		return ErrFull
	}

	r.storage[r.writeIdx] = item
	r.writeIdx = r.advance(r.writeIdx)
	r.full = r.writeIdx == r.readIdx

	return nil
}

func (r *RingBuf[T]) Poll() (v T, e error) {
	if r.isEmpty() {
		return v, ErrEmpty
	}

	i := r.readIdx

	r.readIdx = r.advance(r.readIdx)
	r.full = false
	return r.storage[i], nil
}

func (r *RingBuf[T]) isEmpty() bool {
	return r.writeIdx == r.readIdx && (!r.full)
}

func (r *RingBuf[T]) isFull() bool {
	return r.full
}

func (r *RingBuf[T]) advance(i uint64) uint64 {
	return (i + 1) % uint64(len(r.storage))
}

func (r *RingBuf[T]) String() string {
	return fmt.Sprintf(" %+v", r.storage)
}

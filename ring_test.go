package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type A struct {
	a int
	b int
}

func TestNew(t *testing.T) {
	_, err := New[int8](0)
	assert.Error(t, err, "New should return error")

	b, err := New[int8](2)
	assert.NoError(t, err, "New should not return error")
	assert.NotNil(t, b, "ring buffer should be not nil")
}

func TestOfferPoll(t *testing.T) {
	b, _ := New[int](3)

	_, err := b.Poll()

	assert.Error(t, err, "poll should return error on empty buffer")

	err = b.Offer(10)
	assert.NoError(t, err, "should be able to offer")

	err = b.Offer(11)
	assert.NoError(t, err, "should be able to offer")

	err = b.Offer(12)
	assert.NoError(t, err, "should be able to offer")

	err = b.Offer(13)
	assert.Error(t, err, "offer should fail")

	fmt.Println(b)

	v, err := b.Poll()
	if assert.NoError(t, err, "poll should not return error") {
		assert.Equal(t, 10, v, "poll should return first value offered to buffer")
	}

	v, err = b.Poll()
	if assert.NoError(t, err, "poll should not return error") {
		assert.Equal(t, 11, v, "poll should return first value offered to buffer")
	}

	v, err = b.Poll()
	if assert.NoError(t, err, "poll should not return error") {
		assert.Equal(t, 12, v, "poll should return first value offered to buffer")
	}

	_, err = b.Poll()
	assert.Error(t, err, "poll should return error")
}

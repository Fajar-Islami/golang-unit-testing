package stack_test

import (
	"testing"

	"github.com/Fajar-Islami/golang-testing/stack"
	"github.com/stretchr/testify/assert"
)

// Requirement
// * A Stack is empty on construction
// * A Stack has size 0 on construction
// * After n pushes to an empty stack (n>0), the stack is non-empty and its size equals n
// * If one pushes x then pops, the value pooped is x, and the size is one less than old size
// * If one pushes x then peeks, the value returned is x, but the size stays the same
// * If the is in n, then after pops, the stack is empty and has size 0
// * Popping form an empty stack return an error: ErrNoSuchElement
// * Peeking form an empty stack return an error: ErrNoSuchElement

func TestNewSet(t *testing.T) {
	t.Run("A Stack is empty on construction", func(t *testing.T) {
		s := stack.New()
		assert.True(t, s.IsEmpty())
	})

	t.Run("A Stack has size 0 on construction", func(t *testing.T) {
		s := stack.New()
		assert.Equal(t, 0, s.Size())
	})
}

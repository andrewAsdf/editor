package main

import (
	"editor/core"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBuffer(t *testing.T) {
	t.Run("Read_Empty", func(t *testing.T) {
		buf := core.NewEmptyBuffer()
		result := buf.Read(0, 0)
		assert.Equal(t, "", string(result))
	})

	t.Run("Insert_Then_Read", func(t *testing.T) {
		buf := core.NewEmptyBuffer()
		buf.Insert([]rune("asdf"), 0, 0)
		result := buf.Read(0, 4)
		assert.Equal(t, "asdf", string(result))
	})

	t.Run("Read_FromPosition", func(t *testing.T) {
		buf := core.NewBuffer(strings.NewReader("0123456789"))
		result := buf.Read(1, 4)
		assert.Equal(t, "123", string(result))
	})

	t.Run("Insert_ToPosition", func(t *testing.T) {
		buf := core.NewBuffer(strings.NewReader("0123456789"))
		buf.Insert([]rune("asdf"), 0, 4)
		result := buf.Read(0, 10)
		assert.Equal(t, "asdf456789", string(result))
	})
}
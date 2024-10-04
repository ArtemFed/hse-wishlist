package xcommon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	studentKey = "studentKey"
	teacherKey = "teacherKey"
	adminKey   = "adminKey"
	unknownKey = "unknownKey"
)

func TestSet(t *testing.T) {
	t.Parallel()

	t.Run("Test adding and removing elements", func(t *testing.T) {
		t.Parallel()

		set := make(Set)

		set.Add(studentKey)
		set.Add(teacherKey)
		set.Add(adminKey)

		assert.True(t, set.Contains(studentKey))
		assert.True(t, set.Contains(teacherKey))
		assert.True(t, set.Contains(adminKey))
		assert.False(t, set.Contains(unknownKey))

		set.Remove(teacherKey)

		assert.False(t, set.Contains(teacherKey))
		assert.Equal(t, 2, set.Size())
	})

	t.Run("Test getting items", func(t *testing.T) {
		t.Parallel()

		set := make(Set)
		set.Add(studentKey)
		set.Add(teacherKey)
		set.Add(adminKey)

		items := set.Items()

		assert.Len(t, items, 3)
		assert.Contains(t, items, studentKey)
		assert.Contains(t, items, teacherKey)
		assert.Contains(t, items, adminKey)
	})

	t.Run("Test nil set", func(t *testing.T) {
		t.Parallel()

		set := make(Set)
		set = nil

		assert.False(t, set.Contains(studentKey))
	})
}

func TestSet_AddItems(t *testing.T) {
	t.Parallel()

	t.Run("Test adding items", func(t *testing.T) {
		t.Parallel()

		set := make(Set)

		set.AddItems([]string{studentKey, teacherKey, adminKey})

		assert.True(t, set.Contains(studentKey))
		assert.True(t, set.Contains(teacherKey))
		assert.True(t, set.Contains(adminKey))
		assert.Equal(t, 3, set.Size())
	})

	t.Run("Test adding items to an existing set", func(t *testing.T) {
		t.Parallel()

		set := make(Set)
		set.Add(studentKey)

		set.AddItems([]string{teacherKey, adminKey})

		assert.True(t, set.Contains(studentKey))
		assert.True(t, set.Contains(teacherKey))
		assert.True(t, set.Contains(adminKey))
		assert.Equal(t, 3, set.Size())
	})
}

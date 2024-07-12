package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCacheLRU(t *testing.T) {

	lruCache := NewCacheLRU(1)

	t.Run("should not find key", func(t *testing.T) {
		t.Parallel()
		value := lruCache.Get("a")

		assert.Equal(t, -1, value)
	})

	t.Run("should add key on list", func(t *testing.T) {
		t.Parallel()

		lruCache.Set("b", 10)
		val := lruCache.Get("b")

		assert.Equal(t, 10, val)
	})

	t.Run("should remove the first element", func(t *testing.T) {
		t.Parallel()

		lruCache.Set("a", 1)
		val := lruCache.Get("b")

		assert.Equal(t, -1, val)
	})
}

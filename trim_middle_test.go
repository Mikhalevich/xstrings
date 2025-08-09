package xstrings_test

import (
	"testing"

	"github.com/Mikhalevich/xstrings"

	"github.com/stretchr/testify/require"
)

func TestTrimMiddle(t *testing.T) {
	t.Parallel()

	t.Run("Hello world => ... (-1)", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", -1)
		require.Equal(t, "...", actual)
	})

	t.Run("Hello world => ... (0)", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 0)
		require.Equal(t, "...", actual)
	})

	t.Run("Hello world => ... (1)", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 1)
		require.Equal(t, "...", actual)
	})

	t.Run("Hello world => ... (2)", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 2)
		require.Equal(t, "...", actual)
	})

	t.Run("Hello world => ... (3)", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 3)
		require.Equal(t, "...", actual)
	})

	t.Run("Hello world => H...", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 4)
		require.Equal(t, "H...", actual)
	})

	t.Run("Hello world => H...d", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 5)
		require.Equal(t, "H...d", actual)
	})

	t.Run("Hello world => He...d", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 6)
		require.Equal(t, "He...d", actual)
	})

	t.Run("Hello world => He...ld", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 7)
		require.Equal(t, "He...ld", actual)
	})

	t.Run("Hello world => Hel...ld", func(t *testing.T) {
		t.Parallel()

		actual := xstrings.TrimMiddle("Hello world", 8)
		require.Equal(t, "Hel...ld", actual)
	})
}

func BenchmarkTrimMiddle(b *testing.B) {
	for b.Loop() {
		xstrings.TrimMiddle("Hello world", 5)
	}
}

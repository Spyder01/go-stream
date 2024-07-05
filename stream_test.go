package stream

import (
	"strings"
	"testing"
)

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestStreamInt(t *testing.T) {
	list := []int{9, 1, 2, 4, 5, 6, 7}
	stream := Stream[int](list)

	stream.
		Filter(func(ele int) bool { return ele%2 == 0 }).
		Map(func(ele int) int { return ele * ele })

	result := stream.Collect()

	if !Equal(result, []int{4, 16, 36}) {
		t.Error(result, " !== ", []int{4, 16, 36})
	}
}

func TestStreamString(t *testing.T) {
	list := []string{"akduycikuadyfc", "jygdxujyasdfgdyjarsfdua", "ytugdiuasgdisegdujhbhjagvjadfjdf"}

	result := Stream[string](list).
		Filter(func(ele string) bool {
			return !strings.HasPrefix(ele, "a")
		}).
		Collect()

	if len(result) != 2 {
		t.Error("Test Failed.")
	}
}

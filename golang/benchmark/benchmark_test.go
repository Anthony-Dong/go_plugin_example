package benchmark

import "testing"

func BenchmarkGoPlugin(b *testing.B) {
	add := LoadAddFunc("../../plugin/output/plugin.so")
	for i := 0; i < b.N; i++ {
		if x := add(1, 2); x != 3 {
			b.Fatal(`err !!!`)
		}
	}
}

func BenchmarkGoNative(b *testing.B) {
	add := func(x, y int) int {
		return x + y
	}
	for i := 0; i < b.N; i++ {
		if x := add(1, 2); x != 3 {
			b.Fatal(`err !!!`)
		}
	}
}

func TestName(t *testing.T) {
	add := LoadAddFunc("../plugin/output/plugin.so")
	t.Log(add(1, 2))
}

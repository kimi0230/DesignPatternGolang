package singleton

import "testing"

func TestSingleton(t *testing.T) {
	single1 := GetInstance()
	single2 := GetInstance()

	if single1 != single2 {
		t.Error("")
	}
}

func TestSingletonOnce(t *testing.T) {
	single1 := GetInstanceOnce()
	single2 := GetInstanceOnce()

	if single1 != single2 {
		t.Error("")
	}
}

func BenchmarkSingleton(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetInstance()
	}
}
func BenchmarkSingletonOnce(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetInstanceOnce()
	}
}

/*
go test -benchmem -run=none DesignPatternGolang/Singleton/ -bench=.
BenchmarkSingleton-8            425317406                3.655 ns/op           0 B/op          0 allocs/op
BenchmarkSingletonOnce-8        679671385                1.857 ns/op           0 B/op          0 allocs/op
*/

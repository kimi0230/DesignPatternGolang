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
goos: darwin
goarch: amd64
pkg: DesignPatternGolang/Singleton
cpu: Intel(R) Core(TM) i5-6400 CPU @ 2.70GHz
BenchmarkSingleton-4            401843040                2.971 ns/op           0 B/op          0 allocs/op
BenchmarkSingletonOnce-4        581861808                2.129 ns/op           0 B/op          0 allocs/op
PASS
ok      DesignPatternGolang/Singleton   2.960s
*/

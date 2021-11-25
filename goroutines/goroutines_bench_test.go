package goroutines

import (
	"io"
	"math"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

type rd struct {
	r   *rand.Rand
	max int64
	n   int64
}

func NewReader(n int64) *rd {
	return &rd{
		r:   rand.New(rand.NewSource(time.Now().UnixNano())),
		max: n,
	}
}

func (r *rd) Read(p []byte) (n int, err error) {
	if atomic.LoadInt64(&r.n) >= r.max {
		return 0, nil
	}

	for k := range p {
		p[k] = byte(r.r.Int())
	}

	atomic.AddInt64(&r.n, 1)

	return len(p), nil
}

func run(b *testing.B, x uint, y uint, n int64, f func(reader io.Reader, n int64)) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		reader := NewReader(int64(math.Pow(float64(x), float64(y))))
		b.StartTimer()
		f(reader, n)
		b.StopTimer()
	}
}

func Benchmark_oneTimeRoutines10x3_8go(b *testing.B) {
	run(b, 10, 3, 8, oneTimeRoutines)
}

func Benchmark_longRoutines10x3_8go(b *testing.B) {
	run(b, 10, 3, 8, longRoutines)
}

func Benchmark_oneTimeRoutines10x3_16go(b *testing.B) {
	run(b, 10, 3, 16, oneTimeRoutines)
}

func Benchmark_longRoutines10x3_16go(b *testing.B) {
	run(b, 10, 3, 16, longRoutines)
}

func Benchmark_oneTimeRoutines10x3_24go(b *testing.B) {
	run(b, 10, 3, 24, oneTimeRoutines)
}

func Benchmark_longRoutines10x3_24go(b *testing.B) {
	run(b, 10, 3, 24, longRoutines)
}

func Benchmark_oneTimeRoutines10x3_48go(b *testing.B) {
	run(b, 10, 3, 48, oneTimeRoutines)
}

func Benchmark_longRoutines10x3_48go(b *testing.B) {
	run(b, 10, 3, 48, longRoutines)
}

func Benchmark_oneTimeRoutines10x3_96go(b *testing.B) {
	run(b, 10, 3, 96, oneTimeRoutines)
}

func Benchmark_longRoutines10x3_96go(b *testing.B) {
	run(b, 10, 3, 96, longRoutines)
}

func Benchmark_oneTimeRoutines10x7_8go(b *testing.B) {
	run(b, 10, 7, 8, oneTimeRoutines)
}

func Benchmark_longRoutines10x7_8go(b *testing.B) {
	run(b, 10, 7, 8, longRoutines)
}

func Benchmark_oneTimeRoutines10x7_16go(b *testing.B) {
	run(b, 10, 7, 16, oneTimeRoutines)
}

func Benchmark_longRoutines10x7_16go(b *testing.B) {
	run(b, 10, 7, 16, longRoutines)
}

func Benchmark_oneTimeRoutines10x7_24go(b *testing.B) {
	run(b, 10, 7, 24, oneTimeRoutines)
}

func Benchmark_longRoutines10x7_24go(b *testing.B) {
	run(b, 10, 7, 24, longRoutines)
}

func Benchmark_oneTimeRoutines10x7_48go(b *testing.B) {
	run(b, 10, 7, 48, oneTimeRoutines)
}

func Benchmark_longRoutines10x7_48go(b *testing.B) {
	run(b, 10, 7, 48, longRoutines)
}

func Benchmark_oneTimeRoutines10x7_96go(b *testing.B) {
	run(b, 10, 7, 96, oneTimeRoutines)
}

func Benchmark_longRoutines10x7_96go(b *testing.B) {
	run(b, 10, 7, 96, longRoutines)
}

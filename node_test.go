package quadtree

import (
	"math/rand"
	"testing"
	"time"
)

var (
	boxFullSize  = NewRectangle(-360.0, 360.0, -90, 90)
	boxHundredKM = NewRectangle(10.0, 10.9, 42.0, 42.9)
	boxTenKM     = NewRectangle(10.0, 10.09, 42.0, 42.09)
)

func benchmarkSearch(amount int, box *Rectangle, b *testing.B) {
	// generate random seed
	rand.Seed(time.Now().UnixNano())

	// create QuadTree
	qt := NewQuadTree()

	// populate grid with random points
	for i := 0; i < amount; i++ {
		x := (rand.Float32() * 360.0) - 180.0
		y := (rand.Float32() * 180.0) - 90.0

		qt.Insert(
			NewPoint(x, y, i),
		)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		qt.Search(box)
	}
}

func Benchmark_Search_FullSize_100(b *testing.B)    { benchmarkSearch(100, boxFullSize, b) }
func Benchmark_Search_FullSize_1000(b *testing.B)   { benchmarkSearch(1000, boxFullSize, b) }
func Benchmark_Search_FullSize_10000(b *testing.B)  { benchmarkSearch(10000, boxFullSize, b) }
func Benchmark_Search_FullSize_50000(b *testing.B)  { benchmarkSearch(50000, boxFullSize, b) }
func Benchmark_Search_FullSize_100000(b *testing.B) { benchmarkSearch(100000, boxFullSize, b) }
func Benchmark_Search_FullSize_500000(b *testing.B) { benchmarkSearch(500000, boxFullSize, b) }

func Benchmark_Search_100_Km_100(b *testing.B)    { benchmarkSearch(100, boxHundredKM, b) }
func Benchmark_Search_100_Km_1000(b *testing.B)   { benchmarkSearch(1000, boxHundredKM, b) }
func Benchmark_Search_100_Km_10000(b *testing.B)  { benchmarkSearch(10000, boxHundredKM, b) }
func Benchmark_Search_100_Km_50000(b *testing.B)  { benchmarkSearch(50000, boxHundredKM, b) }
func Benchmark_Search_100_Km_100000(b *testing.B) { benchmarkSearch(100000, boxHundredKM, b) }
func Benchmark_Search_100_Km_500000(b *testing.B) { benchmarkSearch(500000, boxHundredKM, b) }

func Benchmark_Search_10_Km_100(b *testing.B)    { benchmarkSearch(100, boxTenKM, b) }
func Benchmark_Search_10_Km_1000(b *testing.B)   { benchmarkSearch(1000, boxTenKM, b) }
func Benchmark_Search_10_Km_10000(b *testing.B)  { benchmarkSearch(10000, boxTenKM, b) }
func Benchmark_Search_10_Km_50000(b *testing.B)  { benchmarkSearch(50000, boxTenKM, b) }
func Benchmark_Search_10_Km_100000(b *testing.B) { benchmarkSearch(100000, boxTenKM, b) }
func Benchmark_Search_10_Km_500000(b *testing.B) { benchmarkSearch(500000, boxTenKM, b) }

func Benchmark_Update(b *testing.B) {
	// create QuadTree
	qt := NewQuadTree()

	// create start point
	p := NewPoint(0, 0, "test")

	// Reset timers
	b.ReportAllocs()
	b.ResetTimer()

	// Loop
	for n := 0; n < b.N; n++ {
		// create new random position point
		x := (rand.Float32() * 360.0) - 180.0
		y := (rand.Float32() * 180.0) - 90.0

		np := NewPoint(x, y, "test")

		// Update
		qt.Update(p, np)

		// Save old point
		p = np
	}
}

func Benchmark_Insert(b *testing.B) {
	// create QuadTree
	qt := NewQuadTree()

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		x := (rand.Float32() * 360.0) - 180.0
		y := (rand.Float32() * 180.0) - 90.0

		qt.Insert(
			NewPoint(x, y, b.N),
		)
	}
}

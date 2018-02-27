# quadtree
QuadTree implementation in Go

## Install
`go get github.com/mitjaziv/quadtree`

## Benchmarks
Some benchmarks on my machine, with different amount of data and search region sizes.

```
Benchmark_Search_FullSize_100-8      	  200000	      9532 ns/op
Benchmark_Search_FullSize_1000-8     	   10000	    111103 ns/op
Benchmark_Search_FullSize_10000-8    	    2000	    778906 ns/op
Benchmark_Search_FullSize_50000-8    	     200	   6368964 ns/op
Benchmark_Search_FullSize_100000-8   	     100	  18387614 ns/op
Benchmark_Search_FullSize_500000-8   	      10	 101077023 ns/op
Benchmark_Search_100_Km_100-8        	10000000	       180 ns/op
Benchmark_Search_100_Km_1000-8       	 3000000	       399 ns/op
Benchmark_Search_100_Km_10000-8      	 1000000	      1485 ns/op
Benchmark_Search_100_Km_50000-8      	  300000	      4157 ns/op
Benchmark_Search_100_Km_100000-8     	  200000	      9490 ns/op
Benchmark_Search_100_Km_500000-8     	   20000	     73385 ns/op
Benchmark_Search_10_Km_100-8         	10000000	       214 ns/op
Benchmark_Search_10_Km_1000-8        	 5000000	       290 ns/op
Benchmark_Search_10_Km_10000-8       	 3000000	       469 ns/op
Benchmark_Search_10_Km_50000-8       	 1000000	      1302 ns/op
Benchmark_Search_10_Km_100000-8      	  500000	      3172 ns/op
Benchmark_Search_10_Km_500000-8      	  100000	     21258 ns/op
Benchmark_Update-8                   	20000000	       100 ns/op	      32 B/op	       1 allocs/op
Benchmark_Insert-8                   	 3000000	       439 ns/op	      73 B/op	       2 allocs/op
```

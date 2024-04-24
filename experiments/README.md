# Experiments

Different performance experiments

Tested on `go1.21.0 windows/amd64`

## Sum

Given an array of vectors and a vector A, return a new array that is the result of adding A to each element of the original array.

For small array sizes, the mutable vector outperforms the immutable vector. But as the vector size grows, the opposite happens. Immutable vectors begin to out perform.

```
goos: windows
goarch: amd64
pkg: github.com/EliCDavis/vector/experiments
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkSumVector/Len-100-20   	 9706915	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-1000-20  	 1000000	      1133 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-10000-20 	   89824	     13145 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-100000-20         	   10000	    148563 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-1000000-20        	     511	   2662477 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-10000000-20       	      38	  26449755 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-100-20     	 9983551	       134.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-1000-20    	 1000000	      1113 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-10000-20   	   98847	     10837 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-100000-20  	    9955	    136423 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-1000000-20 	     666	   2055645 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-10000000-20         	      42	  27580652 ns/op	       0 B/op	       0 allocs/op
```

Formatted in CSV

```csv
vector type, array length, times ran, ns/op, B/op, allocs/op
Vector, 100,                 9706915,     116.1,  0, 0
Vector, 1000,                1000000,      1133,  0, 0
Vector, 10000,                 89824,     13145,  0, 0
Vector, 100000,                10000,    148563,  0, 0
Vector, 1000000,                 511,   2662477,  0, 0
Vector, 10000000,                 38,  26449755,  0, 0
Mutable Vector, 100,         9983551,     134.0,  0, 0
Mutable Vector, 1000,        1000000,      1113,  0, 0
Mutable Vector, 10000,         98847,     10837,  0, 0
Mutable Vector, 100000,         9955,    136423,  0, 0
Mutable Vector, 1000000,         666,   2055645,  0, 0
Mutable Vector, 10000000,         42,  27580652,  0, 0
```
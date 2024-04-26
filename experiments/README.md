# Experiments

Different performance experiments to help inform this libraries design.

Golang: `go1.21.0 windows/amd64`

CPU: `13th Gen Intel(R) Core(TM) i7-13800H`

## Sum

Given an array of vectors and a vector A, return a new array that is the result of adding A to each element of the original array.

### Results

For small array sizes, the immutable vector outperforms the mutable vector. But as the vector size grows, the opposite happens. Mutable vectors begin to out perform. Running this code on non x86 machine brings different, more predictable results however.

### Take Away

The benchmarking results is probably just noise at the moment. This is further proved out in the `Add` benchmark.

### Raw Data

```
goos: windows
goarch: amd64
pkg: github.com/EliCDavis/vector/experiments
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkSumVector/Len-100-20   	         9706915	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-1000-20  	         1000000	      1133 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-10000-20 	           89824	     13145 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-100000-20         	   10000	    148563 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-1000000-20        	     511	   2662477 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumVector/Len-10000000-20       	      38	  26449755 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-100-20     	 9983551	       134.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-1000-20    	 1000000	      1113 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-10000-20   	   98847	     10837 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-100000-20  	    9955	    136423 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-1000000-20 	     666	   2055645 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumMutableVector/Len-10000000-20         42	  27580652 ns/op	       0 B/op	       0 allocs/op
```

Formatted in CSV

```csv
vector type,    arr len,   times ran,     ns/op,    B/op, allocs/op
Vector,         100,         9706915,     116.1,    0,    0
Vector,         1000,        1000000,      1133,    0,    0
Vector,         10000,         89824,     13145,    0,    0
Vector,         100000,        10000,    148563,    0,    0
Vector,         1000000,         511,   2662477,    0,    0
Vector,         10000000,         38,  26449755,    0,    0
Mutable Vector, 100,         9983551,     134.0,    0,    0
Mutable Vector, 1000,        1000000,      1113,    0,    0
Mutable Vector, 10000,         98847,     10837,    0,    0
Mutable Vector, 100000,         9955,    136423,    0,    0
Mutable Vector, 1000000,         666,   2055645,    0,    0
Mutable Vector, 10000000,         42,  27580652,    0,    0
```

## Add

Add one vector to another, saving the results away

### Results

Both methods in which the structs, immutable or mutable, who take a struct as an argument and returns a new struct as the result, outperform everything else. Anytime a pointer is involved, performance starts to degrade. Side effects required for modifying the value of a mutable vector seems to degrade performance.

### Take Away

Avoid using pointers. Having a method return a value isn't bad.

### Raw Data

```
goos: windows
goarch: amd64
pkg: github.com/EliCDavis/vector/experiments
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkAddVector-20                                           	1000000000	         0.8899 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVector-20                                    	1000000000	         0.8704 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVectorInPlace-20                             	372655183	         3.291 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVectorInPlaceUsingPointer-20                 	396180162	         3.012 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVectorInPlaceWithReturn-20                   	350640970	         3.527 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVectorInPlaceWithReturnUsingPointer-20       	396528655	         3.077 ns/op	       0 B/op	       0 allocs/op
BenchmarkAddMutableVectorInPlaceTakingPointerUsingPointer-20    	393230407	         3.249 ns/op	       0 B/op	       0 allocs/op
```

Formatted in CSV

```csv
benchmark,                                                  times ran,   ns/op,   B/op,  allocs/op
Add Vector,                                                 1000000000,  0.8899,  0,     0
Add Mutable Vector,                                         1000000000,  0.8704,  0,     0
Add Mutable Vector In Place,                                372655183,   3.291,   0,     0
Add Mutable Vector In Place Using Pointer,                  396180162,   3.012,   0,     0
Add Mutable Vector In Place With Return,                    350640970,   3.527,   0,     0
Add Mutable Vector In Place With Return Using Pointer,      396528655,   3.077,   0,     0
Add Mutable Vector In Place Taking Pointer Using Pointer,   393230407,   3.249,   0,     0
```

## Accessor

Testing whether or not using a function to access a specific component in a vector is detrimental to perforamnce.

### Results

Both using a function to access a variable of a struct, and just referencing the variable directly perform almost the same. Given it's sub nanosecond, any difference is hard to tease out due to it probably being noise.

### Takeaway

Using a function to access a component in a vector is Okay.

### Raw Data

```
goos: windows
goarch: amd64
pkg: github.com/EliCDavis/vector/experiments
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkAccessVector-20           1000000000         0.5130 ns/op       0 B/op       0 allocs/op
BenchmarkAccessMutableVector-20    1000000000         0.5163 ns/op       0 B/op       0 allocs/op
```

Formatted in CSV

```csv
benchmark,              times ran,   ns/op,     B/op,  allocs/op
Access Vector           1000000000,  0.5130,    0,     0,
Access MutableVector    1000000000,  0.5163,    0,     0,
```
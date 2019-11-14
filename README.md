# Vector

Collection of **immutable** vector math functions I've written overtime for different hobby projects. I don't really think I'll change method names, just adding new ones over time as I see fit.

Has support for both Vector2 (x, y) and Vector3 (x, y, z) functions.

```
go get github.com/EliCDavis/vector
```

## Example

```go
// Move all points by a certain amount
func Translate(starting []vector.Vector3, amt vector.Vector3) []vector.Vector3 {
	results := make([]vector.Vector3, len(starting))

	for i := 0; i < len(starting); i++ {
		results[i] = starting[i].Add(amt)
	}

	return results
}
```

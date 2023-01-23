# Vector

Collection of **generic, immutable** vector math functions I've written overtime for different hobby projects. 

Has support for both Vector2 (x, y) and Vector3 (x, y, z) functions.

## Example

```go
// Move all points by a certain amount
func Translate[T vector.Number](starting []vector3.Vector[T], amount vector3.Vector[T]) []vector3.Vector[T] {
	results := make([]vector3.Vector[T], len(starting))

	for i := 0; i < len(starting); i++ {
		results[i] = starting[i].Add(amt)
	}

	return results
}
```

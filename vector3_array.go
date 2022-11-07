package vector

type Vector3Array []Vector3

func (v3a Vector3Array) Add(other Vector3) (out Vector3Array) {
	out = make(Vector3Array, len(v3a))

	for i, v := range v3a {
		out[i] = v.Add(other)
	}

	return
}

func (v3a Vector3Array) MultByConstant(t float64) (out Vector3Array) {
	out = make(Vector3Array, len(v3a))

	for i, v := range v3a {
		out[i] = v.MultByConstant(t)
	}

	return
}

func (v3a Vector3Array) DivByConstant(t float64) (out Vector3Array) {
	out = make(Vector3Array, len(v3a))

	for i, v := range v3a {
		out[i] = v.DivByConstant(t)
	}

	return
}

func (v3a Vector3Array) Normalized() (out Vector3Array) {
	out = make(Vector3Array, len(v3a))

	for i, v := range v3a {
		out[i] = v.Normalized()
	}

	return
}

func (v3a Vector3Array) Sum() (sum Vector3) {
	for _, v := range v3a {
		sum = sum.Add(v)
	}
	return
}

func (v3a Vector3Array) Modify(f func(Vector3) Vector3) (out Vector3Array) {
	out = make(Vector3Array, len(v3a))

	for i, v := range v3a {
		out[i] = f(v)
	}

	return
}

// Average sums all vector3's components together and divides each
// component by the number of values added
func (v3a Vector3Array) Average(vectors []Vector3) Vector3 {
	return v3a.Sum().DivByConstant(float64(len(vectors)))
}

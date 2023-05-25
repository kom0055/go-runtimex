package generics

func Pointer[T any](v T) *T {
	return &v
}

func SlicePointer[T any](v []T) []*T {
	res := make([]*T, 0, len(v))
	for i := range v {
		vi := v[i]
		res = append(res, &vi)
	}
	return res
}

func SliceValue[T any](v []*T) []T {
	res := make([]T, 0, len(v))
	for i := range v {
		vi := v[i]
		if vi == nil {
			continue
		}
		res = append(res, *vi)
	}
	return res
}

func MapValuePointer[K comparable, V any](m map[K]V) map[K]*V {
	res := map[K]*V{}
	for k := range m {
		v := m[k]
		res[k] = &v
	}
	return res
}

func MapValueValue[K comparable, V any](m map[K]*V) map[K]V {
	res := map[K]V{}
	for k := range m {
		v := m[k]
		if v == nil {
			continue
		}
		res[k] = *v
	}
	return res
}

func Value[T any](p *T) T {
	if p == nil {
		return Default[T]()
	}
	return *p
}

func Default[T any]() T {
	var defaultVal T
	return defaultVal
}

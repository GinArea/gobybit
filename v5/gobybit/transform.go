package gobybit

func forward[T any](v T) (T, error) {
	return v, nil
}

func transformList[I, T any](l []I, transform func(I) (T, error)) (v []T, err error) {
	v = make([]T, len(l))
	for n, s := range l {
		v[n], err = transform(s)
		if err != nil {
			break
		}
	}
	return
}

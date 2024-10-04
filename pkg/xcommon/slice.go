package xcommon

import "errors"

func EnsureSingle[T any](sl []T) (*T, error) {
	if len(sl) == 0 {
		return nil, errors.New("slice is empty")
	}
	if len(sl) > 1 {
		return nil, errors.New("slice contains more than 1 item")
	}
	return &sl[0], nil
}

func ConvertSlice[F any, T any](sl []F, converter func(item F) T) []T {
	newSl := make([]T, len(sl))
	for i, v := range sl {
		newSl[i] = converter(v)
	}
	return newSl
}

func ConvertSliceP[F any, T any](sl []F, converter func(item *F) *T) []T {
	newSl := make([]T, len(sl))
	for i, v := range sl {
		newSl[i] = *converter(&v)
	}
	return newSl
}

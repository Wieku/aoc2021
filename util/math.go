package util

import "golang.org/x/exp/constraints"

func Abs[T constraints.Integer | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}

	return a
}

func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Clamp[T constraints.Integer | constraints.Float](x, min, max T) T {
	return Min(max, Max(min, x))
}

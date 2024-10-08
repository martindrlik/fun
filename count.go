package fun

import "iter"

func Count[T any](it iter.Seq[T]) (count int) {
	for range it {
		count++
	}
	return
}

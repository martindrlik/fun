package fun

// Dup returns v, v for given v.
func Dup[V any](v V) (V, V) { return v, v }

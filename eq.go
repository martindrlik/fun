package fun

func EqAlwaysTrue[V1 any, V2 any](V1, V2) bool { return true }
func Eq[V comparable](v1, v2 V) bool           { return v1 == v2 }

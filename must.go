package fun

// Must panics if there is an error; otherwise returns t.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

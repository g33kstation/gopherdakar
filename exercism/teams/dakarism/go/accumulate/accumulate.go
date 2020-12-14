package accumulate

// Converter is a function that do things
type Converter func(string) string

// Accumulate perform an operation on each element a collection
func Accumulate(given []string, f Converter) []string {
	trans := []string{}
	for _, shim := range given {
		trans = append(trans, f(shim))
	}
	return trans
}

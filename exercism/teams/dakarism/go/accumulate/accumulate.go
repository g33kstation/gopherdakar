package accumulate

// Converter is a function that do things
type Converter func(string) string

// Accumulate perform an operation on each element a collection
func Accumulate(given []string, converter Converter) []string {
	return []string{}
}

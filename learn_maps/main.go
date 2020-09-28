// Gophers Dakar - 2020
package main

import "fmt"

var temperatures = map[string]map[string]temperature{}

type temperature struct {
	min, max float32
}

func (t temperature) String() string {
	return fmt.Sprintf("(%2.1f, %2.1f)/C", t.min, t.max)
}

func main() {
	temperatures["dakar"] = make(map[string]temperature)

	temperatures["dakar"]["jan"] = temperature{
		min: 12.2,
		max: 29.65,
	}
	fmt.Printf("temperatures:\n\t%v\n", temperatures)
}

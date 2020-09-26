// Gophers Dakar - 2020
package main

import "fmt"

type temperature struct {
	min, max float32
}

type ftemperature struct {
	min, max float32
}

func (t ftemperature) String() string {
	return fmt.Sprintf("(%2.1f, %2.1f)/F", t.min, t.max)
}

func (t temperature) String() string {
	return fmt.Sprintf("(%2.1f, %2.1f)/C", t.min, t.max)
}

func (t temperature) toFareinheight() ftemperature {
	temp := ftemperature{}
	temp.min = t.min * -17.22222
	temp.max = t.max * -17.22222
	return temp
}

func main() {
	temps := make(map[string]map[string]temperature)

	dakar := make(map[string]temperature)
	dakar["jan"] = temperature{20.9, 22.8}
	dakar["feb"] = temperature{26.9, 22.8}
	dakar["mar"] = temperature{21.9, 22.8}
	dakar["avr"] = temperature{21.9, 22.8}
	dakar["mai"] = temperature{18.9, 22.8}
	dakar["jun"] = temperature{21.9, 22.8}
	dakar["jui"] = temperature{21.9, 22.8}
	dakar["aou"] = temperature{25.9, 22.8}
	dakar["sep"] = temperature{17.9, 26.8}
	dakar["nov"] = temperature{19.9, 22.8}
	dakar["dec"] = temperature{21.9, 22.8}
	temps["dakar"] = dakar

	for region, temperatures := range temps {
		fmt.Printf("%s\n", region)
		for mois, temperature := range temperatures {
			converted := temperature.toFareinheight()
			fmt.Printf("\t%s %s %s \n", mois, temperature, converted)
		}
	}
}

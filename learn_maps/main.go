// Gophers Dakar - 2020
package main

import "fmt"

type person struct {
	nom, prenom string
	age         int
}

func (p person) String() string {
	return fmt.Sprintf("L'utilisateur est %s %s, et il a %d ans!\n", p.prenom, p.nom, p.age)
}

func main() {
	Omar := person{"FAll", "Omar", 42}
	fmt.Println(Omar)

	temp := make(map[string]float32)
	temp["Dakar"] = 35.2
	temp["Thiès"] = 40.1
	temp["Matam"] = 42.3

	for i, t := range temp {
		fmt.Printf("%s : %f\n", i, t)
	}
}

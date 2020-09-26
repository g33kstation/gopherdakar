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
	// Go est compatible utf-8 par default (caratère arabe, chinois, emojis)
	fmt.Println("Hello world 😋")

	fmt.Println(Omar)
}

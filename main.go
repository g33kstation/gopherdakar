// Gophers Dakar - 2020
package main

import "fmt"

// declaration de variables globales
// suivi
var prenom string = "Issa" // ""
var age int = 23           // 0
var budget float64         // 0.0 (zeroed value)
var bogoss bool            // false (zeroed value)

//déclaration d'une structure
type person struct {
	nom, prenom string
	age         int
}

func main() {

	Omar := person{"FAll", "Omar", 42}
	// Go est compatible utf-8 par default (caratère arabe, chinois, emojis)
	fmt.Println("Hello world 😋")

	fmt.Printf("%#V\n", prenom)
	fmt.Printf("%#V\n", age)

	fmt.Printf("%#V\n", bogoss)
	fmt.Printf("%#V\n", Omar)
}

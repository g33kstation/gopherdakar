// Gophers Dakar - 2020
package main

import (
	"fmt"
	"strings"
)

type temperature struct {
	min, max float32
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

type ftemperature struct {
	min, max float32
}

func (t ftemperature) String() string {
	return fmt.Sprintf("(%2.1f, %2.1f)/F", t.min, t.max)
}

type region struct {
	nom          string
	temperatures [12]temperature
}

func (r region) String() string {
	return fmt.Sprintf("Région: %s\n%s", r.nom, r.temperaturesString())
}

var mois = [12]string{"jan", "fév", "mar", "avr", "mai", "jun", "jui", "aoû", "sep", "oct", "nov", "dec"}

func (r region) temperaturesString() string {
	str := []string{}
	for idx, t := range r.temperatures {
		mois := mois[idx]
		converted := t.toFareinheight()
		str = append(str, fmt.Sprintf("  %s %s %s", mois, t.String(), converted.String()))
	}
	return strings.Join(str, "\n")
}

func main() {
	dk := region{"Dakar", [12]temperature{}}
	fmt.Println(dk)
}

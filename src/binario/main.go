package main

import (
	"fmt"
)

type Valor uint8
type Casilla struct {
	valor       Valor
	descubierta bool
}
type Tablero [][]Casilla

func main() {
	vacio := Casilla{2, false}
	cero := Casilla{0, true}
	uno := Casilla{1, true}
	var tablero Tablero = [][]Casilla{
		{uno, vacio, vacio, vacio, vacio, vacio},
		{uno, uno, vacio, vacio, uno, vacio},
		{cero, uno, uno, vacio, cero, uno},
		{vacio, vacio, vacio, vacio, vacio, uno},
		{cero, vacio, uno, vacio, uno, vacio},
		{cero, cero, vacio, vacio, cero, uno},
	}

	tablero.Imprimir()
}

func (tablero Tablero) Imprimir() {
	for i := 0; i < len(tablero); i++ {
		for j := 0; j < len(tablero[i]); j++ {
			fmt.Print("|")

			casilla := tablero[i][j]
			if casilla.descubierta {
				fmt.Print(casilla.valor)
			} else {
				fmt.Print(" ")
			}

			if j == len(tablero[i])-1 {
				fmt.Print("|")
			}
		}
		fmt.Println("")
	}
}

package main

import (
	"fmt"
)

type Valor uint8
type Tablero [][]Valor

func main() {
	tablero := make(Tablero, 0)
	cantidadColumnas := 10
	cantidadFilas := 10

	for i := 0; i < cantidadColumnas; i++ {
		columna := make([]Valor, 0)
		for j := 0; j < cantidadFilas; j++ {
			columna = append(columna, 1)
		}

		tablero = append(tablero, columna)

	}

	tablero.Imprimir()
}

func (tablero Tablero) Imprimir() {
	for i := 0; i < len(tablero); i++ {
		for j := 0; j < len(tablero[i]); j++ {
			fmt.Printf("|%v|", tablero[i][j])
		}
		fmt.Println("")
	}
}

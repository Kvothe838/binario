package main

import (
	"fmt"
)

type Valor uint8
type Casilla struct {
	valor   Valor
	visible bool
}
type Tablero [][]Casilla

func main() {
	tablero := ArmarTablero()
	fmt.Println("Tablero inicial")
	fmt.Println()
	tablero.Imprimir()

	fmt.Println()
	fmt.Println()

	tableroArmado := tablero.Resolver()
	fmt.Println("Tablero resuelto")
	fmt.Println()
	tableroArmado.Imprimir()
}

func ArmarTablero() Tablero {
	vacio := Casilla{2, false}
	cero := Casilla{0, true}
	uno := Casilla{1, true}
	var tablero Tablero = [][]Casilla{
		{vacio, uno, uno, vacio, uno, vacio, vacio, vacio, vacio, vacio, vacio, vacio, uno, vacio},
		{vacio, vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio, vacio, cero, vacio, vacio},
		{uno, vacio, vacio, vacio, cero, cero, vacio, cero, cero, vacio, uno, vacio, vacio, vacio},
		{vacio, cero, cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, uno},
		{vacio, cero, vacio, vacio, vacio, cero, vacio, vacio, cero, vacio, vacio, vacio, vacio, vacio},
		{vacio, vacio, vacio, vacio, vacio, cero, vacio, vacio, vacio, vacio, uno, uno, vacio, vacio},
		{cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio},
		{vacio, cero, vacio, vacio, uno, vacio, cero, vacio, cero, vacio, vacio, cero, vacio, vacio},
		{uno, vacio, vacio, vacio, vacio, vacio, vacio, vacio, cero, vacio, vacio, vacio, uno, vacio},
		{vacio, vacio, uno, uno, vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio, vacio},
		{vacio, cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, uno},
		{uno, vacio, vacio, cero, vacio, uno, vacio, vacio, cero, vacio, vacio, vacio, vacio, uno},
		{vacio, vacio, vacio, vacio, vacio, vacio, cero, vacio, cero, cero, vacio, vacio, vacio, vacio},
		{vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio},
	}

	return tablero
}

func (tablero Tablero) Imprimir() {
	for i := 0; i < len(tablero); i++ {
		for j := 0; j < len(tablero[i]); j++ {
			fmt.Print("|")

			casilla := tablero[i][j]
			if casilla.visible {
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

func (tablero Tablero) Resolver() Tablero {
	tablero = tablero.ResolverDoblesSeguidos()
	return tablero
}

func (tablero Tablero) ResolverDoblesSeguidos() Tablero {
	seguirBarriendo := false

	for !seguirBarriendo {
		seguirBarriendo = false

		for indiceFila := 0; indiceFila < len(tablero); indiceFila++ {
			fila := tablero[indiceFila]

			for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
				casilla := fila[indiceColumna]
				primerColumna := 0
				ultimaColumna := len(fila) - 1
				siguienteColumna := indiceColumna + 1
				anteriorColumna := indiceColumna - 1
				opuesto := casilla.valor.ObtenerOpuesto()

				if !casilla.visible || indiceColumna == ultimaColumna {
					continue
				}

				siguienteCasilla := fila[siguienteColumna]

				if !siguienteCasilla.visible {
					continue
				}

				if casilla.valor != siguienteCasilla.valor {
					continue
				}

				if indiceColumna != primerColumna {
					anteriorCasilla := &fila[anteriorColumna]

					if !anteriorCasilla.visible {
						anteriorCasilla.valor = opuesto
						anteriorCasilla.visible = true
						seguirBarriendo = true
					}

				}

				if siguienteColumna != ultimaColumna {
					siguienteSiguienteCasilla := &fila[siguienteColumna+1]

					if !siguienteCasilla.visible {
						siguienteSiguienteCasilla.valor = opuesto
						siguienteSiguienteCasilla.visible = true
						seguirBarriendo = true
					}
				}
			}
		}
	}

	return tablero
}

func (valor Valor) ObtenerOpuesto() Valor {
	opuestos := map[Valor]Valor{
		0: 1,
		1: 0,
		2: 2,
	}

	return opuestos[valor]
}

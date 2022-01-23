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
type Girable func(*bool) Tablero

func main() {
	tablero := ArmarTablero()
	fmt.Println("Tablero inicial")
	fmt.Println()
	tablero.Imprimir()

	fmt.Println()
	fmt.Println()

	/* tableroArmado := tablero.Resolver() */
	tablero.Resolver()

	/* fmt.Println()
	fmt.Println()

	fmt.Println("Tablero resuelto")
	fmt.Println()
	tableroArmado.Imprimir() */
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
	volverABarrer := true
	nuevoTablero := tablero

	for volverABarrer {
		volverABarrer = false
		nuevoTablero.ResolverDoblesSeguidos()
		fmt.Println("Dobles seguidos")
		fmt.Println()
		nuevoTablero.Imprimir()

		fmt.Println()
		fmt.Println()

		nuevoTablero.ResolverDoblesSalteados(&volverABarrer)
		fmt.Println("Dobles salteados")
		fmt.Println()
		nuevoTablero.Imprimir()

		fmt.Println()
		fmt.Println()

		nuevoTablero.ResolverFaltaUnNumero(&volverABarrer)
		fmt.Println("Falta un numero")
		fmt.Println()
		nuevoTablero.Imprimir()
	}

	return nuevoTablero
}

func (tablero *Tablero) ResolverDoblesSeguidos() {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverDoblesSeguidosHorizontal(&seguirBarriendo)
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverDoblesSeguidosHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
			}
		}

		tablero.DarVuelta()
	}
}

func (tablero *Tablero) ResolverDoblesSalteados(volverABarrerExterno *bool) {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverDoblesSalteadosHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				*volverABarrerExterno = true
			}
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverDoblesSalteadosHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
				*volverABarrerExterno = true
			}
		}

		tablero.DarVuelta()
	}
}

func (tablero *Tablero) ResolverFaltaUnNumero(volverABarrerExterno *bool) {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverFaltaUnNumeroHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				*volverABarrerExterno = true
			}
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverFaltaUnNumeroHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
				*volverABarrerExterno = true
			}
		}

		tablero.DarVuelta()
	}
}

func (tablero Tablero) ResolverDoblesSeguidosHorizontal(seguirBarriendo *bool) Tablero {
	*seguirBarriendo = false

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
					*seguirBarriendo = true
				}

			}

			if siguienteColumna != ultimaColumna {
				siguienteSiguienteCasilla := &fila[siguienteColumna+1]

				if !siguienteSiguienteCasilla.visible {
					siguienteSiguienteCasilla.valor = opuesto
					siguienteSiguienteCasilla.visible = true
					*seguirBarriendo = true
				}
			}
		}
	}

	return tablero
}

func (tablero Tablero) ResolverDoblesSalteadosHorizontal(seguirBarriendo *bool) Tablero {
	*seguirBarriendo = false

	for indiceFila := 0; indiceFila < len(tablero); indiceFila++ {
		fila := tablero[indiceFila]

		for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
			casilla := fila[indiceColumna]
			ultimaColumna := len(fila) - 1
			siguienteColumna := indiceColumna + 1
			opuesto := casilla.valor.ObtenerOpuesto()

			if !casilla.visible || indiceColumna == ultimaColumna || siguienteColumna == ultimaColumna {
				continue
			}

			siguienteCasilla := &fila[siguienteColumna]

			if siguienteCasilla.visible {
				continue
			}

			siguienteSiguienteCasilla := fila[siguienteColumna+1]

			if !siguienteSiguienteCasilla.visible || casilla.valor != siguienteSiguienteCasilla.valor {
				continue
			}

			siguienteCasilla.valor = opuesto
			siguienteCasilla.visible = true
			*seguirBarriendo = true
		}
	}

	return tablero
}

func (tablero Tablero) ResolverFaltaUnNumeroHorizontal(seguirBarriendo *bool) Tablero {
	*seguirBarriendo = false
	lado := len(tablero)
	vecesMaximoEnFila := lado / 2

	for indiceFila := 0; indiceFila < lado; indiceFila++ {
		fila := tablero[indiceFila]
		vecesEnFilaPorValor := map[Valor]int{
			0: 0,
			1: 0,
			2: 0,
		}

		hayAlMenosUnoNoVisible := false

		for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
			casilla := fila[indiceColumna]
			vecesEnFilaPorValor[casilla.valor]++

			if !casilla.visible {
				hayAlMenosUnoNoVisible = true
			}
		}

		if hayAlMenosUnoNoVisible {
			var valorACompletar *Valor = nil
			cantidadCeros := vecesEnFilaPorValor[0]
			cantidadUnos := vecesEnFilaPorValor[1]

			if cantidadCeros == vecesMaximoEnFila {
				uno := Valor(1)
				valorACompletar = &uno
			} else if cantidadUnos == vecesMaximoEnFila {
				cero := Valor(0)
				valorACompletar = &cero
			}

			if valorACompletar != nil {
				for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
					casilla := &fila[indiceColumna]

					if !casilla.visible {
						casilla.valor = *valorACompletar
						casilla.visible = true
					}
				}

				*seguirBarriendo = true
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

func (tablero *Tablero) DarVuelta() {
	lado := len(*tablero)
	nuevoTablero := make(Tablero, lado)
	filas := len(*tablero)

	for i := 0; i < filas; i++ {
		nuevoTablero[i] = make([]Casilla, lado)
	}

	for i := 0; i < filas; i++ {
		columnas := len((*tablero)[i])

		for j := 0; j < columnas; j++ {
			nuevoTablero[j][i] = (*tablero)[i][j]
		}
	}

	*tablero = nuevoTablero
}

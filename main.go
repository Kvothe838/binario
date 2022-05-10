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
type Casillas []Casilla
type Girable func(*bool) Tablero

var imprimirPasos bool

func main() {
	tablero := ArmarTablero()
	fmt.Println("Tablero inicial")
	fmt.Println()
	tablero.Imprimir()

	fmt.Println()
	fmt.Println()
	imprimirPasos = false
	tableroArmado := tablero.Resolver()

	fmt.Println()
	fmt.Println()

	fmt.Println("Tablero resuelto")
	fmt.Println()
	tableroArmado.Imprimir()
}

func ArmarTablero() Tablero {
	vacio := Casilla{2, false}
	cero := Casilla{0, true}
	uno := Casilla{1, true}
	var tablero Tablero = Tablero{
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

	/* var tablero Tablero = Tablero{
		{vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, cero, vacio, vacio, vacio},
		{cero, cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, uno, uno, vacio},
		{vacio, cero, vacio, vacio, vacio, cero, vacio, vacio, vacio, uno, vacio, vacio, vacio, vacio},
		{vacio, vacio, uno, vacio, uno, vacio, vacio, vacio, vacio, vacio, vacio, cero, vacio, vacio},
		{vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio, cero, cero, vacio, vacio, vacio, vacio},
		{vacio, cero, vacio, vacio, vacio, uno, vacio, vacio, vacio, uno, vacio, vacio, cero, vacio},
		{uno, cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio},
		{vacio, vacio, uno, vacio, uno, vacio, vacio, cero, vacio, vacio, vacio, cero, vacio, vacio},
		{vacio, vacio, vacio, vacio, vacio, vacio, uno, vacio, vacio, vacio, uno, vacio, vacio, cero},
		{vacio, vacio, vacio, vacio, cero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio},
		{vacio, cero, vacio, vacio, vacio, uno, vacio, uno, vacio, vacio, vacio, vacio, cero, vacio},
		{vacio, vacio, vacio, vacio, cero, vacio, vacio, vacio, cero, vacio, uno, vacio, vacio, cero},
		{cero, vacio, vacio, uno, vacio, vacio, uno, uno, vacio, uno, vacio, cero, cero, vacio},
		{cero, vacio, vacio, uno, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, cero, vacio},
	} */

	return tablero
}

func (tablero Tablero) Imprimir() {
	for i := 0; i < len(tablero); i++ {
		fila := Casillas(tablero[i])
		fila.Imprimir()
		fmt.Println("")
	}
}

func (casillas Casillas) Imprimir() {
	for indiceCasilla := 0; indiceCasilla < len(casillas); indiceCasilla++ {
		fmt.Print("|")

		casilla := casillas[indiceCasilla]
		casilla.Imprimir()

		if indiceCasilla == len(casillas)-1 {
			fmt.Print("|")
		}
	}
}

func (casillas Casillas) ImprimirTodo() {
	for indiceCasilla := 0; indiceCasilla < len(casillas); indiceCasilla++ {
		fmt.Print("|")

		casilla := casillas[indiceCasilla]
		fmt.Print(casilla.valor)

		if indiceCasilla == len(casillas)-1 {
			fmt.Print("|")
		}
	}
}

func (casilla Casilla) Imprimir() {
	if casilla.visible {
		fmt.Print(casilla.valor)
	} else {
		fmt.Print(" ")
	}
}

func (tablero Tablero) Resolver() Tablero {
	volverABarrer := true
	nuevoTablero := tablero

	for volverABarrer {
		volverABarrer = false

		if imprimirPasos {
			fmt.Println("Dobles seguidos")
			fmt.Println()
		}

		nuevoTablero.ResolverDoblesSeguidos()

		if imprimirPasos {
			fmt.Println()
			fmt.Println()

			nuevoTablero.Imprimir()

			fmt.Println()
			fmt.Println()
		}

		if nuevoTablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Dobles salteados")
			fmt.Println()
		}

		nuevoTablero.ResolverDoblesSalteados(&volverABarrer)

		if imprimirPasos {
			fmt.Println()
			fmt.Println()

			nuevoTablero.Imprimir()

			fmt.Println()
			fmt.Println()
		}

		if nuevoTablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Falta un número")
			fmt.Println()
		}

		nuevoTablero.ResolverFaltaUnNumero(&volverABarrer)

		if imprimirPasos {
			fmt.Println()
			fmt.Println()

			nuevoTablero.Imprimir()

			fmt.Println()
			fmt.Println()
		}

		if nuevoTablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Falta uno de un valor")
			fmt.Println()
		}

		nuevoTablero.ResolverFaltaUnoDeUnValor(&volverABarrer)

		if imprimirPasos {
			fmt.Println()
			fmt.Println()

			nuevoTablero.Imprimir()

			fmt.Println()
			fmt.Println()
		}

		if nuevoTablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Resolver líneas duplicadas")
			fmt.Println()
		}

		nuevoTablero.ResolverLineasDuplicadas(&volverABarrer)

		if imprimirPasos {
			fmt.Println()
			fmt.Println()

			nuevoTablero.Imprimir()

			fmt.Println()
			fmt.Println()
		}

		if nuevoTablero.EstaResuelto() {
			break
		}
	}

	return nuevoTablero
}

func (tablero Tablero) EstaResuelto() bool {
	for indiceFila := 0; indiceFila < len(tablero); indiceFila++ {
		for indiceColumna := 0; indiceColumna < len(tablero[indiceFila]); indiceColumna++ {
			if !tablero[indiceFila][indiceColumna].visible {
				return false
			}
		}
	}

	return true
}

func (tablero *Tablero) ResolverDoblesSeguidos() {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverDoblesSeguidosHorizontal(&seguirBarriendo)
		}

		if tablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Doy vuelta")
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverDoblesSeguidosHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
			}
		}

		if imprimirPasos {
			tablero.DarVueltaAlReves()
			fmt.Println("Vuelvo a la posición original")
		}

		if tablero.EstaResuelto() {
			break
		}
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

		if tablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Doy vuelta")
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

		if imprimirPasos {
			fmt.Println("Vuelvo a la posición original")
		}
		tablero.DarVueltaAlReves()

		if tablero.EstaResuelto() {
			break
		}
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

		if tablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Doy vuelta")
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

		if imprimirPasos {
			fmt.Println("Vuelvo a la posición original")
		}

		tablero.DarVueltaAlReves()

		if tablero.EstaResuelto() {
			break
		}
	}
}

func (tablero *Tablero) ResolverFaltaUnoDeUnValor(volverABarrerExterno *bool) {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverFaltaUnoDeUnValorHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				*volverABarrerExterno = true
			}
		}

		if tablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Doy vuelta")
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverFaltaUnoDeUnValorHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
				*volverABarrerExterno = true
			}
		}

		if imprimirPasos {
			fmt.Println("Vuelvo a la posición original")
		}

		tablero.DarVueltaAlReves()

		if tablero.EstaResuelto() {
			break
		}
	}
}

func (tablero *Tablero) ResolverLineasDuplicadas(volverABarrerExterno *bool) {
	volverABarrer := true

	for volverABarrer {
		volverABarrer = false
		seguirBarriendo := true

		for seguirBarriendo {
			tablero.ResolverLineasDuplicadasHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				*volverABarrerExterno = true
			}
		}

		if tablero.EstaResuelto() {
			break
		}

		if imprimirPasos {
			fmt.Println("Doy vuelta")
		}

		tablero.DarVuelta()

		seguirBarriendo = true
		for seguirBarriendo {
			tablero.ResolverLineasDuplicadasHorizontal(&seguirBarriendo)

			if seguirBarriendo {
				volverABarrer = true
				*volverABarrerExterno = true
			}
		}

		if imprimirPasos {
			fmt.Println("Vuelvo a la posición original")
		}

		tablero.DarVueltaAlReves()

		if tablero.EstaResuelto() {
			break
		}
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
					if imprimirPasos {
						fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, anteriorColumna, opuesto)
						fmt.Println()
					}
					anteriorCasilla.valor = opuesto
					anteriorCasilla.visible = true
					*seguirBarriendo = true
				}
			}

			if siguienteColumna != ultimaColumna {
				siguienteSiguienteCasilla := &fila[siguienteColumna+1]

				if !siguienteSiguienteCasilla.visible {
					if imprimirPasos {
						fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, siguienteColumna+1, opuesto)
						fmt.Println()
					}
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

			if imprimirPasos {
				fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, siguienteColumna, opuesto)
				fmt.Println()
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
		vecesEnFilaPorValor, hayAlMenosUnoNoVisible := tablero.ObtenerVecesEnFilaPorValor(fila)
		/* fmt.Printf("indiceFila: %v | vecesEnFilaPorValor: %+v | hayAlMenosUnoNoVisible: %v", indiceFila, vecesEnFilaPorValor, hayAlMenosUnoNoVisible)
		fmt.Println() */

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

			if valorACompletar == nil {
				continue
			}

			for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
				casilla := &fila[indiceColumna]

				if casilla.visible {
					continue
				}

				if imprimirPasos {
					fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, indiceColumna, *valorACompletar)
					fmt.Println()
				}

				casilla.valor = *valorACompletar
				casilla.visible = true
			}

			*seguirBarriendo = true
		}
	}

	return tablero
}

type VecesEnFilaPorValor map[Valor]int

func (v VecesEnFilaPorValor) String() string {
	stringToReturn := "{"

	for k, v := range v {
		stringToReturn += fmt.Sprintf("valor: %v | veces: %v", k, v)
		stringToReturn += fmt.Sprintln()
	}

	stringToReturn += "}"

	return stringToReturn
}

func (tablero Tablero) ObtenerVecesEnFilaPorValor(fila []Casilla) (map[Valor]int, bool) {
	vecesEnFilaPorValor := map[Valor]int{
		0: 0,
		1: 0,
	}

	hayAlMenosUnoNoVisible := false

	for indiceColumna := 0; indiceColumna < len(fila); indiceColumna++ {
		casilla := fila[indiceColumna]

		if !casilla.visible {
			hayAlMenosUnoNoVisible = true
			continue
		}

		vecesEnFilaPorValor[casilla.valor]++
	}

	return vecesEnFilaPorValor, hayAlMenosUnoNoVisible
}

func (tablero Tablero) ResolverFaltaUnoDeUnValorHorizontal(seguirBarriendo *bool) Tablero {
	*seguirBarriendo = false
	lado := len(tablero)
	vecesMaximoEnFila := lado / 2

	// Se recorre cada fila
	for indiceFila := 0; indiceFila < lado; indiceFila++ {
		fila := tablero[indiceFila]
		// Se obtienen la cantidad de veces que se repite cada valor, y si hay al menos un valor no visible
		vecesEnFilaPorValor, hayAlMenosUnoNoVisible := tablero.ObtenerVecesEnFilaPorValor(fila)
		/* fmt.Printf("indiceFila: %v | vecesEnFilaPorValor: %+v | hayAlMenosUnoNoVisible: %v", indiceFila, vecesEnFilaPorValor, hayAlMenosUnoNoVisible)
		fmt.Println() */

		if !hayAlMenosUnoNoVisible {
			continue
		}

		var valorACompletar *Valor = nil
		cantidadCeros := vecesEnFilaPorValor[0]
		cantidadUnos := vecesEnFilaPorValor[1]

		if cantidadCeros == vecesMaximoEnFila-1 {
			// Si pueden haber como máxmo 7 ceros y hay 6, significa que el algoritmo va a evaluar qué pasa si se llena con cero cada casilla.
			// Si se llega a una contradicción, la casilla se llena de su opuesto, el 1.
			uno := Valor(1)
			valorACompletar = &uno
		} else if cantidadUnos == vecesMaximoEnFila-1 {
			cero := Valor(0)
			valorACompletar = &cero
		}

		if valorACompletar == nil {
			/* fmt.Printf("Rompo con fila %v", indiceFila)
			fmt.Println() */
			continue
		}

		opuesto := (*valorACompletar).ObtenerOpuesto()
		// Se crea una copia de fila para que no se rellenen por error las casillas hipotéticas que usa el algoritmo.
		filaAux := fila
		indicesAACtualizar := make([]int, 0)

		if imprimirPasos {
			fmt.Print("línea principio: ")
			Casillas(filaAux).Imprimir()
			fmt.Println()
		}

		// Se recorre cada columna
		for indiceColumna := 0; indiceColumna < lado; indiceColumna++ {
			casilla := &filaAux[indiceColumna]

			if casilla.visible {
				continue
			}

			// Se llena la casilla del valor opuesto (el valor que le falta una repetición) y luego se llenan todas las demás del valor a completar.
			casilla.valor = opuesto

			for indiceColumnaAux := 0; indiceColumnaAux < lado; indiceColumnaAux++ {
				casillaAux := &filaAux[indiceColumnaAux]

				/* fmt.Println("Paso por casilla: ")
				casillaAux.Imprimir()
				fmt.Println() */

				if indiceColumnaAux == indiceColumna || casillaAux.visible {
					/* fmt.Println("continúo") */
					continue
				}

				/* fmt.Printf("Valor: %v", *valorACompletar) */

				casillaAux.valor = *valorACompletar
			}

			/* fmt.Print("línea intermedia: ")
			Casillas(filaAux).ImprimirTodo()
			fmt.Println() */

			// Se van a contar la cantidad del valor opuesto que existan seguidos en la fila. Si son más de 2, quiere decir que es una combinación
			// imposible. Entonces, si al haber llenado la original el opuesto se genera una combinación imposible, quiere decir que la original va
			// del valor a completar.
			cantidadSeguidos := 1

			/* fmt.Printf("filaAux: %+v", Casillas(filaAux))
			fmt.Println() */
			for indiceColumnaAux := 0; indiceColumnaAux < lado; indiceColumnaAux++ {
				casillaAux := filaAux[indiceColumnaAux]

				// Contar como seguido si el valor es el buscado (el valor a completar) y la columna no es la primera (ya que sino no se podría
				// acceder a [indiceColumnaAux-1]) y el valor anterior es igual al actual.
				if casillaAux.valor == *valorACompletar && indiceColumnaAux != 0 && filaAux[indiceColumnaAux-1].valor == *valorACompletar {
					cantidadSeguidos++
				} else {
					cantidadSeguidos = 1
				}

				// Si ya van más de 2 seguidos, añadir la columna como para actualizar con el valor a completar y salir del loop ya que no hace falta
				// ver si hay más seguidos, la contradicción está cumplida.
				if cantidadSeguidos > 2 {
					indicesAACtualizar = append(indicesAACtualizar, indiceColumna)
					break
				}
			}

			for indiceColumnaAux := 0; indiceColumnaAux < lado; indiceColumnaAux++ {
				casillaAux := &filaAux[indiceColumnaAux]

				if indiceColumnaAux == indiceColumna || casillaAux.visible {
					continue
				}

				casillaAux.valor = 2
			}

			filaAux = fila
			/* fmt.Print("línea reseteada: ")
			Casillas(filaAux).Imprimir()
			fmt.Println() */
		}

		for _, indiceAACtualizar := range indicesAACtualizar {
			/* fmt.Printf("Índice a actualizar: %v", indiceAACtualizar)
			fmt.Println() */
			*seguirBarriendo = true
			/* fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, indiceAACtualizar, *valorACompletar)
			fmt.Println() */
			fila[indiceAACtualizar].valor = *valorACompletar
			fila[indiceAACtualizar].visible = true
		}
	}

	return tablero
}

// Dos líneas no pueden tener los mismos números en las mismas posiciones. Tienen que diferir en, por lo menos, dos celdas.
func (tablero Tablero) ResolverLineasDuplicadasHorizontal(seguirBarriendo *bool) Tablero {
	*seguirBarriendo = false
	lado := len(tablero)

	// Recorro cada fila
	for indiceFila := 0; indiceFila < lado; indiceFila++ {
		// Acá voy a guardar la posición en columna de la celda no visible en la fila original (como máximo, solo pueden diferir en una posición)
		var posicionNoVisibleFila1 *int = nil
		var posicionNoVisibleFila2 *int = nil
		fila := &tablero[indiceFila]

		// Recorro cada una de las demás filas (todas excepto la que ya estoy recorriendo)
		for indiceFilaAux := 0; indiceFilaAux < lado; indiceFilaAux++ {
			filaAux := &tablero[indiceFilaAux]
			// Acá voy a guardar las posiciones en columnas de las celdas no visibles en la fila comparada
			var posicionNoVisibleFilaAux1 *int = nil
			var posicionNoVisibleFilaAux2 *int = nil

			if indiceFila == indiceFilaAux {
				continue
			}

			llenarFila := true

			// Recorro cada columna, ya que voy a comparar dos celdas con la misma posición de columna pero en diferentes filas
			for indiceColumna := 0; indiceColumna < lado; indiceColumna++ {
				celdaOriginal := (*fila)[indiceColumna]
				celdaComparada := (*filaAux)[indiceColumna]

				indiceColumnaSinPuntero := indiceColumna

				// Si ambas celdas son visibles y sus valores son diferentes, entonces la comparación ya no puede realizarse
				if celdaOriginal.visible && celdaComparada.visible && celdaOriginal.valor != celdaComparada.valor {
					llenarFila = false
					break
				}

				if !celdaOriginal.visible || !celdaComparada.visible {
					// Si ambos pares se llenaron, quiere decir que la comparación ya no puede realizarse
					if ObtenerCantidadNil(posicionNoVisibleFila1, posicionNoVisibleFila2,
						posicionNoVisibleFilaAux1, posicionNoVisibleFilaAux2) > 2 {
						llenarFila = false
						break
					}

					if !celdaOriginal.visible {
						// Si la celda original no es visible, entonces hay que ver si es la primera columna en no ser visible, la segunda o la tercera.
						// En caso de ser la priemra va a la fila1, en caso de ser la segunda va a la fila2 y en caso de ser la tercera se rompe el ciclo
						// ya que no se pueden tener más de dos pares de celdas con una no visible
						if posicionNoVisibleFila1 == nil && posicionNoVisibleFilaAux1 == nil {
							posicionNoVisibleFila1 = &indiceColumnaSinPuntero
							continue
						} else if posicionNoVisibleFila2 == nil && posicionNoVisibleFilaAux2 == nil {
							posicionNoVisibleFila2 = &indiceColumnaSinPuntero
							continue
						} else {
							llenarFila = false
							break
						}
					} else if !celdaComparada.visible {
						// Misma lógica que con la original pero se guardan en las aux.
						// Importante que sea un else if, ya que no se puede dar la combinación de dos celdas no visibles para la misma fila, esto no
						// da ninguna información a la hora de deducir valores
						if posicionNoVisibleFila1 == nil && posicionNoVisibleFilaAux1 == nil {
							posicionNoVisibleFilaAux1 = &indiceColumnaSinPuntero
							continue
						} else if posicionNoVisibleFila2 == nil && posicionNoVisibleFilaAux2 == nil {
							posicionNoVisibleFilaAux2 = &indiceColumnaSinPuntero
							continue
						} else {
							llenarFila = false
							break
						}
					}

				}
			}

			if !llenarFila {
				continue
			}

			if posicionNoVisibleFila1 != nil || posicionNoVisibleFilaAux1 != nil {
				*seguirBarriendo = true
				if posicionNoVisibleFila1 != nil {
					fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, *posicionNoVisibleFila1, (*filaAux)[*posicionNoVisibleFila1].valor)
					fmt.Println()
					(*fila)[*posicionNoVisibleFila1].valor = (*filaAux)[*posicionNoVisibleFila1].valor.ObtenerOpuesto()
					(*fila)[*posicionNoVisibleFila1].visible = true
				} else if posicionNoVisibleFilaAux1 != nil {
					fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, *posicionNoVisibleFilaAux1, (*filaAux)[*posicionNoVisibleFilaAux1].valor)
					fmt.Println()
					(*fila)[*posicionNoVisibleFilaAux1].valor = (*filaAux)[*posicionNoVisibleFilaAux1].valor.ObtenerOpuesto()
					(*fila)[*posicionNoVisibleFilaAux1].visible = true
				}
			}

			if posicionNoVisibleFila2 != nil || posicionNoVisibleFilaAux2 != nil {
				*seguirBarriendo = true
				if posicionNoVisibleFila2 != nil {
					fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, *posicionNoVisibleFila2, (*filaAux)[*posicionNoVisibleFila2].valor)
					fmt.Println()
					(*fila)[*posicionNoVisibleFila2].valor = (*filaAux)[*posicionNoVisibleFila2].valor.ObtenerOpuesto()
					(*fila)[*posicionNoVisibleFila2].visible = true
				} else if posicionNoVisibleFilaAux2 != nil {
					fmt.Printf("Lleno casilla (x: %v, y: %v) con %v", indiceFila, *posicionNoVisibleFilaAux2, (*filaAux)[*posicionNoVisibleFilaAux2].valor)
					fmt.Println()
					(*fila)[*posicionNoVisibleFilaAux2].valor = (*filaAux)[*posicionNoVisibleFilaAux2].valor.ObtenerOpuesto()
					(*fila)[*posicionNoVisibleFilaAux2].visible = true
				}
			}
		}
	}

	return tablero
}

func ObtenerCantidadNil(elementos ...*int) int {
	cantidad := 0

	for _, elemento := range elementos {
		if elemento != nil {
			cantidad++
		}
	}

	return cantidad
}

func (valor Valor) ObtenerOpuesto() Valor {
	opuestos := map[Valor]Valor{
		0: 1,
		1: 0,
		2: 2,
	}

	return opuestos[valor]
}

// Da una vuelta de 90° antihorario
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
			nuevoTablero[filas-j-1][i] = (*tablero)[i][j]
		}
	}

	*tablero = nuevoTablero
}

// Da una vuelta de 90° horario
// TODO: Optimizar para que en vez de girar 3 veces antihorario, gire una sola horario
func (tablero *Tablero) DarVueltaAlReves() {
	for i := 0; i < 3; i++ {
		tablero.DarVuelta()
	}
}

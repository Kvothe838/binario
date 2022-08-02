package main

import (
	"fmt"
)

type Value uint8
type Box struct {
	value   Value
	visible bool
}
type Grid []Boxes
type Boxes []Box

var printSteps bool

func main() {
	grid := BuildGrid()
	fmt.Println("Initial grid")
	fmt.Println()
	grid.Print()

	fmt.Println()
	fmt.Println()
	printSteps = true
	solvedGrid := grid.Solve()

	fmt.Println()
	fmt.Println()

	fmt.Println("Solved grid")
	fmt.Println()
	solvedGrid.Print()
}

func BuildGrid() Grid {
	empty := Box{2, false}
	zero := Box{0, true}
	one := Box{1, true}
	var grid Grid = Grid{
		{empty, one, one, empty, one, empty, empty, empty, empty, empty, empty, empty, one, empty},
		{empty, empty, empty, empty, empty, empty, one, empty, empty, empty, empty, zero, empty, empty},
		{one, empty, empty, empty, zero, zero, empty, zero, zero, empty, one, empty, empty, empty},
		{empty, zero, zero, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, one},
		{empty, zero, empty, empty, empty, zero, empty, empty, zero, empty, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, zero, empty, empty, empty, empty, one, one, empty, empty},
		{zero, empty, empty, empty, empty, empty, empty, empty, empty, empty, one, empty, empty, empty},
		{empty, zero, empty, empty, one, empty, zero, empty, zero, empty, empty, zero, empty, empty},
		{one, empty, empty, empty, empty, empty, empty, empty, zero, empty, empty, empty, one, empty},
		{empty, empty, one, one, empty, empty, empty, empty, empty, one, empty, empty, empty, empty},
		{empty, zero, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, empty, one},
		{one, empty, empty, zero, empty, one, empty, empty, zero, empty, empty, empty, empty, one},
		{empty, empty, empty, empty, empty, empty, zero, empty, zero, zero, empty, empty, empty, empty},
		{empty, empty, empty, empty, empty, one, empty, empty, empty, empty, empty, one, empty, empty},
	}

	/* var grid Tablero = Tablero{
		{vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, zero, vacio, vacio, vacio},
		{zero, zero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, one, one, vacio},
		{vacio, zero, vacio, vacio, vacio, zero, vacio, vacio, vacio, one, vacio, vacio, vacio, vacio},
		{vacio, vacio, one, vacio, one, vacio, vacio, vacio, vacio, vacio, vacio, zero, vacio, vacio},
		{vacio, vacio, vacio, vacio, one, vacio, vacio, vacio, zero, zero, vacio, vacio, vacio, vacio},
		{vacio, zero, vacio, vacio, vacio, one, vacio, vacio, vacio, one, vacio, vacio, zero, vacio},
		{one, zero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio},
		{vacio, vacio, one, vacio, one, vacio, vacio, zero, vacio, vacio, vacio, zero, vacio, vacio},
		{vacio, vacio, vacio, vacio, vacio, vacio, one, vacio, vacio, vacio, one, vacio, vacio, zero},
		{vacio, vacio, vacio, vacio, zero, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio},
		{vacio, zero, vacio, vacio, vacio, one, vacio, one, vacio, vacio, vacio, vacio, zero, vacio},
		{vacio, vacio, vacio, vacio, zero, vacio, vacio, vacio, zero, vacio, one, vacio, vacio, zero},
		{zero, vacio, vacio, one, vacio, vacio, one, one, vacio, one, vacio, zero, zero, vacio},
		{zero, vacio, vacio, one, vacio, vacio, vacio, vacio, vacio, vacio, vacio, vacio, zero, vacio},
	} */

	return grid
}

func (grid Grid) Print() {
	for i := 0; i < len(grid); i++ {
		row := Boxes(grid[i])
		row.Print()
		fmt.Println("")
	}
}

func (boxes Boxes) Print() {
	for boxIndex := 0; boxIndex < len(boxes); boxIndex++ {
		fmt.Print("|")

		box := boxes[boxIndex]
		box.Print()

		if boxIndex == len(boxes)-1 {
			fmt.Print("|")
		}
	}
}

func (boxes Boxes) PrintAll() {
	for boxIndex := 0; boxIndex < len(boxes); boxIndex++ {
		fmt.Print("|")

		box := boxes[boxIndex]
		fmt.Print(box.value)

		if boxIndex == len(boxes)-1 {
			fmt.Print("|")
		}
	}
}

func (box Box) Print() {
	if box.visible {
		fmt.Print(box.value)
	} else {
		fmt.Print(" ")
	}
}

func (grid Grid) Solve() Grid {
	iterateAgain := true
	newGrid := grid

	for iterateAgain {
		iterateAgain = false

		if printSteps {
			fmt.Println("Doubles in a row")
			fmt.Println()
		}

		newGrid.SolveDoublesInARow()

		if printSteps {
			fmt.Println()
			fmt.Println()

			newGrid.Print()

			fmt.Println()
			fmt.Println()
		}

		if newGrid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Doubles by turns")
			fmt.Println()
		}

		newGrid.SolveDoublesByTurns(&iterateAgain)

		if printSteps {
			fmt.Println()
			fmt.Println()

			newGrid.Print()

			fmt.Println()
			fmt.Println()
		}

		if newGrid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("One number is missing")
			fmt.Println()
		}

		newGrid.SolveMissingNumber(&iterateAgain)

		if printSteps {
			fmt.Println()
			fmt.Println()

			newGrid.Print()

			fmt.Println()
			fmt.Println()
		}

		if newGrid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("One box with one value is missing")
			fmt.Println()
		}

		newGrid.SolveOneBoxOneValue(&iterateAgain)

		if printSteps {
			fmt.Println()
			fmt.Println()

			newGrid.Print()

			fmt.Println()
			fmt.Println()
		}

		if newGrid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Solve duplicated lines")
			fmt.Println()
		}

		newGrid.SolveDuplicatedLines(&iterateAgain)

		if printSteps {
			fmt.Println()
			fmt.Println()

			newGrid.Print()

			fmt.Println()
			fmt.Println()
		}

		if newGrid.IsSolved() {
			break
		}
	}

	return newGrid
}

func (grid Grid) IsSolved() bool {
	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		for columnIndex := 0; columnIndex < len(grid[rowIndex]); columnIndex++ {
			if !grid[rowIndex][columnIndex].visible {
				return false
			}
		}
	}

	return true
}

func (grid *Grid) SolveDoublesInARow() {
	iterateAgain := true

	for iterateAgain {
		iterateAgain = false
		continueIteration := true

		for continueIteration {
			grid.SolveHorizontalDoublesInARow(&continueIteration)
		}

		if grid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Rotate")
		}

		grid.Rotate()

		continueIteration = true
		for continueIteration {
			grid.SolveHorizontalDoublesInARow(&continueIteration)

			if continueIteration {
				iterateAgain = true
			}
		}

		if printSteps {
			grid.RotateBackwards()
			fmt.Println("Return to original position")
		}

		if grid.IsSolved() {
			break
		}
	}
}

func (grid *Grid) SolveDoublesByTurns(externalIterateAgain *bool) {
	iterateAgain := true

	for iterateAgain {
		iterateAgain = false
		continueIteration := true

		for continueIteration {
			grid.ResolverDoblesSalteadosHorizontal(&continueIteration)

			if continueIteration {
				*externalIterateAgain = true
			}
		}

		if grid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Rotate")
		}

		grid.Rotate()

		continueIteration = true
		for continueIteration {
			grid.ResolverDoblesSalteadosHorizontal(&continueIteration)

			if continueIteration {
				iterateAgain = true
				*externalIterateAgain = true
			}
		}

		if printSteps {
			fmt.Println("Rotate backwards")
		}
		grid.RotateBackwards()

		if grid.IsSolved() {
			break
		}
	}
}

func (grid *Grid) SolveMissingNumber(externalIterateAgain *bool) {
	iterateAgain := true

	for iterateAgain {
		iterateAgain = false
		continueIteration := true

		for continueIteration {
			grid.SolveHorizontalMissingNumber(&continueIteration)

			if continueIteration {
				*externalIterateAgain = true
			}
		}

		if grid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Rotate")
		}

		grid.Rotate()

		continueIteration = true
		for continueIteration {
			grid.SolveHorizontalMissingNumber(&continueIteration)

			if continueIteration {
				iterateAgain = true
				*externalIterateAgain = true
			}
		}

		if printSteps {
			fmt.Println("Rotate backwards")
		}

		grid.RotateBackwards()

		if grid.IsSolved() {
			break
		}
	}
}

func (grid *Grid) SolveOneBoxOneValue(externalIterateAgain *bool) {
	iterateAgain := true

	for iterateAgain {
		iterateAgain = false
		continueIteration := true

		for continueIteration {
			grid.SolveHorizontalOneBoxOneValue(&continueIteration)

			if continueIteration {
				*externalIterateAgain = true
			}
		}

		if grid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Rotate")
		}

		grid.Rotate()

		continueIteration = true
		for continueIteration {
			grid.SolveHorizontalOneBoxOneValue(&continueIteration)

			if continueIteration {
				iterateAgain = true
				*externalIterateAgain = true
			}
		}

		if printSteps {
			fmt.Println("Rotate backwards")
		}

		grid.RotateBackwards()

		if grid.IsSolved() {
			break
		}
	}
}

func (grid *Grid) SolveDuplicatedLines(externalIterateAgain *bool) {
	iterateAgain := true

	for iterateAgain {
		iterateAgain = false
		continueIteration := true

		for continueIteration {
			grid.SolveDuplicatedHorizontalLines(&continueIteration)

			if continueIteration {
				*externalIterateAgain = true
			}
		}

		if grid.IsSolved() {
			break
		}

		if printSteps {
			fmt.Println("Rotate")
		}

		grid.Rotate()

		continueIteration = true
		for continueIteration {
			grid.SolveDuplicatedHorizontalLines(&continueIteration)

			if continueIteration {
				iterateAgain = true
				*externalIterateAgain = true
			}
		}

		if printSteps {
			fmt.Println("Rotate backwards")
		}

		grid.RotateBackwards()

		if grid.IsSolved() {
			break
		}
	}
}

func (grid Grid) SolveHorizontalDoublesInARow(continueIteration *bool) Grid {
	*continueIteration = false

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		row := grid[rowIndex]

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			box := row[columnIndex]
			primerColumna := 0
			lastColumn := len(row) - 1
			nextColumn := columnIndex + 1
			previousColumn := columnIndex - 1
			opposite := box.value.GetOpposite()

			if !box.visible || columnIndex == lastColumn {
				continue
			}

			nextBox := row[nextColumn]

			if !nextBox.visible {
				continue
			}

			if box.value != nextBox.value {
				continue
			}

			if columnIndex != primerColumna {
				previousBox := &row[previousColumn]

				if !previousBox.visible {
					if printSteps {
						fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, previousColumn, opposite)
						fmt.Println()
					}
					previousBox.value = opposite
					previousBox.visible = true
					*continueIteration = true
				}
			}

			if nextColumn != lastColumn {
				nextNextBox := &row[nextColumn+1]

				if !nextNextBox.visible {
					if printSteps {
						fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, nextColumn+1, opposite)
						fmt.Println()
					}
					nextNextBox.value = opposite
					nextNextBox.visible = true
					*continueIteration = true
				}
			}
		}
	}

	return grid
}

func (grid Grid) ResolverDoblesSalteadosHorizontal(continueIteration *bool) Grid {
	*continueIteration = false

	for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
		row := grid[rowIndex]

		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			box := row[columnIndex]
			lastColumn := len(row) - 1
			nextColumn := columnIndex + 1
			opposite := box.value.GetOpposite()

			if !box.visible || columnIndex == lastColumn || nextColumn == lastColumn {
				continue
			}

			nextBox := &row[nextColumn]

			if nextBox.visible {
				continue
			}

			nextNextBox := row[nextColumn+1]

			if !nextNextBox.visible || box.value != nextNextBox.value {
				continue
			}

			if printSteps {
				fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, nextColumn, opposite)
				fmt.Println()
			}
			nextBox.value = opposite
			nextBox.visible = true
			*continueIteration = true
		}
	}

	return grid
}

func (grid Grid) SolveHorizontalMissingNumber(continueIteration *bool) Grid {
	*continueIteration = false
	side := len(grid)
	maxTimesInRow := side / 2

	for rowIndex := 0; rowIndex < side; rowIndex++ {
		row := grid[rowIndex]
		timesInRowByValue, isSomeNotVisible := grid.GetTimesInRowByValue(row)

		if isSomeNotVisible {
			var valueToFill *Value = nil
			zerosAmount := timesInRowByValue[0]
			onesAmount := timesInRowByValue[1]

			if zerosAmount == maxTimesInRow {
				one := Value(1)
				valueToFill = &one
			} else if onesAmount == maxTimesInRow {
				zero := Value(0)
				valueToFill = &zero
			}

			if valueToFill == nil {
				continue
			}

			for columnIndex := 0; columnIndex < len(row); columnIndex++ {
				box := &row[columnIndex]

				if box.visible {
					continue
				}

				if printSteps {
					fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, columnIndex, *valueToFill)
					fmt.Println()
				}

				box.value = *valueToFill
				box.visible = true
			}

			*continueIteration = true
		}
	}

	return grid
}

type TimesInRowByValue map[Value]int

func (v TimesInRowByValue) String() string {
	stringToReturn := "{"

	for k, v := range v {
		stringToReturn += fmt.Sprintf("value: %v | times: %v", k, v)
		stringToReturn += fmt.Sprintln()
	}

	stringToReturn += "}"

	return stringToReturn
}

func (grid Grid) GetTimesInRowByValue(row []Box) (map[Value]int, bool) {
	timesInRowByValue := map[Value]int{
		0: 0,
		1: 0,
	}

	isSomeNotVisible := false

	for columnIndex := 0; columnIndex < len(row); columnIndex++ {
		box := row[columnIndex]

		if !box.visible {
			isSomeNotVisible = true
			continue
		}

		timesInRowByValue[box.value]++
	}

	return timesInRowByValue, isSomeNotVisible
}

func (grid Grid) SolveHorizontalOneBoxOneValue(continueIteration *bool) Grid {
	*continueIteration = false
	side := len(grid)
	maxTimesInRow := side / 2

	for rowIndex := 0; rowIndex < side; rowIndex++ {
		row := grid[rowIndex]
		// Se obtienen la cantidad de veces que se repite cada valor, y si hay al menos un valor no visible
		timesInRowByValue, isSomeNotVisible := grid.GetTimesInRowByValue(row)

		if !isSomeNotVisible {
			continue
		}

		var valueToFill *Value = nil
		zerosAmount := timesInRowByValue[0]
		onesAmount := timesInRowByValue[1]

		if zerosAmount == maxTimesInRow-1 {
			// Si pueden haber como máxmo 7 ceros y hay 6, significa que el algoritmo va a evaluar qué pasa si se llena con cero cada casilla.
			// Si se llega a una contradicción, la casilla se llena de su opuesto, el 1.
			one := Value(1)
			valueToFill = &one
		} else if onesAmount == maxTimesInRow-1 {
			zero := Value(0)
			valueToFill = &zero
		}

		if valueToFill == nil {
			continue
		}

		opposite := (*valueToFill).GetOpposite()
		// Se crea una copia de fila para que no se rellenen por error las casillas hipotéticas que usa el algoritmo
		auxRow := row
		indexesToUpdate := make([]int, 0)

		if printSteps {
			fmt.Print("first line: ")
			Boxes(auxRow).Print()
			fmt.Println()
		}

		// Se recorre cada columna
		for columnIndex := 0; columnIndex < side; columnIndex++ {
			box := &auxRow[columnIndex]

			if box.visible {
				continue
			}

			// Se llena la casilla del valor opuesto (el valor que le falta una repetición) y luego se llenan todas las demás del valor a completar.
			box.value = opposite

			for auxColumnIndex := 0; auxColumnIndex < side; auxColumnIndex++ {
				auxBox := &auxRow[auxColumnIndex]

				if auxColumnIndex == columnIndex || auxBox.visible {
					continue
				}

				auxBox.value = *valueToFill
			}

			// Se van a contar la cantidad del value opposite que existan seguidos en la fila. Si son más de 2, quiere decir que es una combinación
			// imposible. Entonces, si al haber llenado la original el opuesto se genera una combinación imposible, quiere decir que la original va
			// del value a completar.
			amountInARow := 1

			for auxColumnIndex := 0; auxColumnIndex < side; auxColumnIndex++ {
				auxBox := auxRow[auxColumnIndex]

				// Contar como seguido si el value es el buscado (el value a completar) y la columna no es la primera (ya que sino no se podría
				// acceder a [auxColumnIndex-1]) y el valor anterior es igual al actual.
				if auxBox.value == *valueToFill && auxColumnIndex != 0 && auxRow[auxColumnIndex-1].value == *valueToFill {
					amountInARow++
				} else {
					amountInARow = 1
				}

				// Si ya van más de 2 seguidos, añadir la columna como para actualizar con el valor a completar y salir del loop ya que no hace falta
				// ver si hay más seguidos, la contradicción está cumplida.
				if amountInARow > 2 {
					indexesToUpdate = append(indexesToUpdate, columnIndex)
					break
				}
			}

			for auxColumnIndex := 0; auxColumnIndex < side; auxColumnIndex++ {
				auxBox := &auxRow[auxColumnIndex]

				if auxColumnIndex == columnIndex || auxBox.visible {
					continue
				}

				auxBox.value = 2
			}

			auxRow = row
		}

		for _, indexToUpdate := range indexesToUpdate {
			*continueIteration = true
			row[indexToUpdate].value = *valueToFill
			row[indexToUpdate].visible = true
		}
	}

	return grid
}

// Dos líneas no pueden tener los mismos números en las mismas posiciones. Tienen que diferir en, por lo menos, dos celdas.
func (grid Grid) SolveDuplicatedHorizontalLines(continueIteration *bool) Grid {
	*continueIteration = false
	side := len(grid)

	for rowIndex := 0; rowIndex < side; rowIndex++ {
		// Acá voy a guardar la posición en columna de la celda no visible en la fila original (como máximo, solo pueden diferir en una posición)
		var notVisiblePositionRow1 *int = nil
		var notVisiblePositionRow2 *int = nil
		row := &grid[rowIndex]

		// Recorro cada una de las demás filas (todas excepto la que ya estoy recorriendo)
		for auxRowIndex := 0; auxRowIndex < side; auxRowIndex++ {
			filaAux := &grid[auxRowIndex]
			// Acá voy a guardar las posiciones en columnas de las celdas no visibles en la fila comparada
			var notVisiblePositionAuxRow1 *int = nil
			var notVisiblePositionAuxRow2 *int = nil

			if rowIndex == auxRowIndex {
				continue
			}

			fillRow := true

			// Recorro cada columna, ya que voy a comparar dos celdas con la misma posición de columna pero en diferentes filas
			for columnIndex := 0; columnIndex < side; columnIndex++ {
				originalBox := (*row)[columnIndex]
				comparingBox := (*filaAux)[columnIndex]

				columnWithoutPointerIndex := columnIndex

				// Si ambas celdas son visibles y sus valores son diferentes, entonces la comparación ya no puede realizarse
				if originalBox.visible && comparingBox.visible && originalBox.value != comparingBox.value {
					fillRow = false
					break
				}

				if !originalBox.visible || !comparingBox.visible {
					// Si ambos pares se llenaron, quiere decir que la comparación ya no puede realizarse
					if GetNilAmount(notVisiblePositionRow1, notVisiblePositionRow2,
						notVisiblePositionAuxRow1, notVisiblePositionAuxRow2) > 2 {
						fillRow = false
						break
					}

					if !originalBox.visible {
						// Si la celda original no es visible, entonces hay que ver si es la primera columna en no ser visible, la segunda o la tercera.
						// En caso de ser la primera va a la fila1, en caso de ser la segunda va a la fila2 y en caso de ser la tercera se rompe el ciclo
						// ya que no se pueden tener más de dos pares de celdas con una no visible
						if notVisiblePositionRow1 == nil && notVisiblePositionAuxRow1 == nil {
							notVisiblePositionRow1 = &columnWithoutPointerIndex
							continue
						} else if notVisiblePositionRow2 == nil && notVisiblePositionAuxRow2 == nil {
							notVisiblePositionRow2 = &columnWithoutPointerIndex
							continue
						} else {
							fillRow = false
							break
						}
					} else if !comparingBox.visible {
						// Misma lógica que con la original pero se guardan en las aux.
						// Importante que sea un else if, ya que no se puede dar la combinación de dos celdas no visibles para la misma fila, esto no
						// da ninguna información a la hora de deducir valores
						if notVisiblePositionRow1 == nil && notVisiblePositionAuxRow1 == nil {
							notVisiblePositionAuxRow1 = &columnWithoutPointerIndex
							continue
						} else if notVisiblePositionRow2 == nil && notVisiblePositionAuxRow2 == nil {
							notVisiblePositionAuxRow2 = &columnWithoutPointerIndex
							continue
						} else {
							fillRow = false
							break
						}
					}
				}
			}

			if !fillRow {
				continue
			}

			if notVisiblePositionRow1 != nil || notVisiblePositionAuxRow1 != nil {
				*continueIteration = true
				if notVisiblePositionRow1 != nil {
					fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, *notVisiblePositionRow1, (*filaAux)[*notVisiblePositionRow1].value)
					fmt.Println()
					(*row)[*notVisiblePositionRow1].value = (*filaAux)[*notVisiblePositionRow1].value.GetOpposite()
					(*row)[*notVisiblePositionRow1].visible = true
				} else if notVisiblePositionAuxRow1 != nil {
					fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, *notVisiblePositionAuxRow1, (*filaAux)[*notVisiblePositionAuxRow1].value)
					fmt.Println()
					(*row)[*notVisiblePositionAuxRow1].value = (*filaAux)[*notVisiblePositionAuxRow1].value.GetOpposite()
					(*row)[*notVisiblePositionAuxRow1].visible = true
				}
			}

			if notVisiblePositionRow2 != nil || notVisiblePositionAuxRow2 != nil {
				*continueIteration = true
				if notVisiblePositionRow2 != nil {
					fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, *notVisiblePositionRow2, (*filaAux)[*notVisiblePositionRow2].value)
					fmt.Println()
					(*row)[*notVisiblePositionRow2].value = (*filaAux)[*notVisiblePositionRow2].value.GetOpposite()
					(*row)[*notVisiblePositionRow2].visible = true
				} else if notVisiblePositionAuxRow2 != nil {
					fmt.Printf("Fill box (x: %v, y: %v) with %v", rowIndex, *notVisiblePositionAuxRow2, (*filaAux)[*notVisiblePositionAuxRow2].value)
					fmt.Println()
					(*row)[*notVisiblePositionAuxRow2].value = (*filaAux)[*notVisiblePositionAuxRow2].value.GetOpposite()
					(*row)[*notVisiblePositionAuxRow2].visible = true
				}
			}
		}
	}

	return grid
}

func GetNilAmount(items ...*int) int {
	amount := 0

	for _, item := range items {
		if item != nil {
			amount++
		}
	}

	return amount
}

func (value Value) GetOpposite() Value {
	opposites := map[Value]Value{
		0: 1,
		1: 0,
		2: 2,
	}

	return opposites[value]
}

// Da una vuelta de 90° antihorario
func (grid *Grid) Rotate() {
	side := len(*grid)
	newGrid := make(Grid, side)
	rows := len(*grid)

	for i := 0; i < rows; i++ {
		newGrid[i] = make([]Box, side)
	}

	for i := 0; i < rows; i++ {
		columns := len((*grid)[i])

		for j := 0; j < columns; j++ {
			newGrid[rows-j-1][i] = (*grid)[i][j]
		}
	}

	*grid = newGrid
}

// Da una vuelta de 90° horario
// TODO: Optimizar para que en vez de girar 3 times antihorario, gire una sola horario
func (grid *Grid) RotateBackwards() {
	for i := 0; i < 3; i++ {
		grid.Rotate()
	}
}

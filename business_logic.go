package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

//GameBusinessLogic is the logic of Game and this part where the magic happens!!
type GameBusinessLogic struct {
	table [][]int
	moves int
}

//LeftMove is the action, when the player type ArrowLeft
func (gameLogic *GameBusinessLogic) LeftMove() [][]int {
	gameLogic.moves++
	newTable := gameLogic.table
	for rowIndex, row := range newTable {
		newTable[rowIndex] = leftRemoveAllWhiteSpace(newTable[rowIndex])
		for cellIndex, cell := range row {
			nextIndex := cellIndex + 1
			if len(row) > nextIndex {
				for i := cellIndex + 1; i < len(row); i++ {
					if cell == row[i] {
						newTable[rowIndex][cellIndex] = cell + row[i]
						newTable[rowIndex][i] = 0
						if len(row)-1 > i {

							newTable[rowIndex] = toLeft(i, newTable[rowIndex])
						}
						break
					}
					if row[i] != 0 || cell != 0 {
						break
					}
				}
			}

		}
	}
	return newTable
}
func leftRemoveAllWhiteSpace(row []int) []int {
	act := 0
	for i := 0; i < len(row)-1; i++ {
		if row[i] == 0 && act < len(row) {
			act++
			row = toLeft(i, row)
			i--
		}
	}
	return row
}
func toLeft(index int, loopRow []int) []int {
	nextIndex := index + 1
	rowWithoutWhiteSpaces := loopRow
	if nextIndex < len(loopRow) {
		loopRow[index] = loopRow[nextIndex]
		return toLeft(nextIndex, loopRow)
	}
	rowWithoutWhiteSpaces[len(rowWithoutWhiteSpaces)-1] = 0
	return rowWithoutWhiteSpaces

}

//RightMove is the action, when the player type ArrowRight
func (gameLogic *GameBusinessLogic) RightMove() [][]int {
	gameLogic.moves++
	newTable := gameLogic.table

	for rowIndex, row := range newTable {
		newTable[rowIndex] = rightRemoveAllWhiteSpace(newTable[rowIndex])
		for cellIndex := len(row) - 1; cellIndex >= 0; cellIndex-- {
			nextIndex := cellIndex - 1
			if nextIndex >= 0 {
				for i := cellIndex - 1; i >= 0; i-- {
					if row[cellIndex] == row[i] {
						newTable[rowIndex][cellIndex] = row[cellIndex] + row[i]
						newTable[rowIndex][i] = 0
						if 0 <= i {
							newTable[rowIndex] = toRight(i, newTable[rowIndex])
						}
						break
					}
					if row[i] != 0 || row[cellIndex] != 0 {
						break
					}
				}
			}
		}
	}
	return newTable
}
func rightRemoveAllWhiteSpace(row []int) []int {
	newRow := row
	act := len(row)
	for i := len(row) - 1; i >= 0; i-- {
		if row[i] == 0 && act > 0 {
			act--
			newRow = toRight(i, newRow)
			i++
		}
	}

	return newRow
}
func toRight(index int, loopRow []int) []int {
	previousIndex := index - 1
	rowWithoutWhiteSpaces := loopRow
	if previousIndex >= 0 {
		rowWithoutWhiteSpaces[index] = rowWithoutWhiteSpaces[previousIndex]
		return toRight(previousIndex, rowWithoutWhiteSpaces)
	}
	rowWithoutWhiteSpaces[0] = 0
	return rowWithoutWhiteSpaces
}

//UpMove is the action, when the player type ArrowUp
func (gameLogic *GameBusinessLogic) UpMove() [][]int {
	gameLogic.table = rotate90(gameLogic.table)
	gameLogic.table = rotate90(gameLogic.LeftMove())
	return gameLogic.table
}

//DownMove is the action, when the player type ArrowDown
func (gameLogic *GameBusinessLogic) DownMove() [][]int {
	gameLogic.table = rotate90(gameLogic.table)
	gameLogic.table = rotate90(gameLogic.RightMove())
	return gameLogic.table
}

func rotate90(table [][]int) [][]int {
	newTable := make([][]int, len(table))
	for indexRow, row := range table {
		if len(row) == 0 {
			newTable[indexRow] = make([]int, len(table))
		}

		for indexCell := range row {
			if len(newTable[indexCell]) == 0 {
				newTable[indexCell] = make([]int, len(table))
			}
			newTable[indexCell][indexRow] = row[indexCell]

		}

	}
	return newTable
}

func template(gameNumbsTable [][]int) []string {
	tables := make([]string, 4, 4)
	for indexTab, gameNumbsRow := range gameNumbsTable {
		var row string
		for _, gameNumbRow := range gameNumbsRow {
			switch {
			case gameNumbRow < 10:
				row = fmt.Sprint(row, "  ", gameNumbRow, " ")
			case gameNumbRow < 100:
				row = fmt.Sprint(row, " ", gameNumbRow, " ")
			case gameNumbRow < 1000:
				row = fmt.Sprint(row, "", gameNumbRow, " ")
			case gameNumbRow < 10000:
				row = fmt.Sprint(row, "", gameNumbRow, "")
			}
		}
		tables[indexTab] = row
	}
	return tables
}

func draw(table []string) {
	term.Clear(term.ColorDefault, term.ColorDefault)
	for indexTable, tab := range table {
		for indexTab, let := range []rune(tab) {
			term.SetCell(indexTab, indexTable, let, term.ColorBlue, term.ColorDefault)
		}
	}
	_ = term.Flush()
}

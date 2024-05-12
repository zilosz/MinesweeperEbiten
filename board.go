package main

import (
	"math/rand"
)

type Board struct {
	Cells [][]*Cell
}

func (board *Board) Cell(coord CellCoordinate) *Cell {
	return board.Cells[coord.Row][coord.Col]
}

func (board *Board) CollectCells() []*Cell {
	cells := make([]*Cell, 0, BoardWidth*BoardHeight)

	for i := 0; i < BoardHeight; i++ {

		for j := 0; j < BoardWidth; j++ {
			cells = append(cells, board.Cells[i][j])
		}
	}

	return cells
}

func (board *Board) CellNeighbors(coord CellCoordinate) []*Cell {
	cells := make([]*Cell, 0, 8)

	for _, neighborPos := range coord.SurroundingNeighbors() {

		if neighborPos.Valid() {
			cells = append(cells, board.Cell(neighborPos))
		}
	}

	return cells
}

func CreateBoard() *Board {
	board := &Board{Cells: make([][]*Cell, BoardWidth*BoardHeight)}

	tempCells := make([]*Cell, BoardWidth*BoardHeight)
	cellIndex := 0

	for i := 0; i < BoardHeight; i++ {
		board.Cells[i] = make([]*Cell, BoardWidth)

		for j := 0; j < BoardWidth; j++ {
			cell := &Cell{Pos: CellCoordinate{Col: j, Row: i}}
			board.Cells[i][j] = cell

			tempCells[cellIndex] = cell
			cellIndex++
		}
	}

	for i := 0; i < BombCount; i++ {
		bombIndex := rand.Intn(len(tempCells))
		tempCells[bombIndex].HasBomb = true
		tempCells = append(tempCells[:bombIndex], tempCells[bombIndex+1:]...)
	}

	for _, cell := range board.CollectCells() {

		if cell.HasBomb {

			for _, neighborCell := range board.CellNeighbors(cell.Pos) {
				neighborCell.BombsNearby++
			}
		}
	}

	return board
}

func (board *Board) RevealBombs() {
	for _, cell := range board.CollectCells() {
		if cell.HasBomb {
			cell.Revealed = true
		}
	}
}

func (board *Board) RevealCell(cellCoord CellCoordinate, revealedCells map[*Cell]bool) bool {
	cell := board.Cell(cellCoord)

	if revealedCells[cell] {
		return false
	}

	cell.Revealed = true
	revealedCells[cell] = true

	if cell.HasBomb {
		board.RevealBombs()
		return true
	}

	if cell.BombsNearby == 0 {

		for _, neighborPos := range cellCoord.AdjacentNeighbors() {

			if neighborPos.Valid() {
				board.RevealCell(neighborPos, revealedCells)
			}
		}
	}

	return false
}

func (board *Board) CheckVictory() bool {
	flaggedBombs := 0
	revealedCount := 0

	for _, cell := range board.CollectCells() {

		if cell.HasBomb {

			if cell.Flagged {
				flaggedBombs++
			}

		} else if cell.Revealed {
			revealedCount++
		}
	}

	return flaggedBombs+revealedCount == BoardWidth*BoardHeight
}

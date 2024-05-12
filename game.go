package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Board           *Board
	LastFocusedCell *Cell
	Finished        bool
	Victory         bool
}

func (g *Game) Reset() {
	g.Board = CreateBoard()
	g.Finished = false
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, cell := range g.Board.CollectCells() {
		cell.Draw(screen)
	}
}

func (g *Game) Layout(int, int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset()
		return nil
	}

	if g.Finished {
		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()
	cellCoord := ScreenCoordinate{X: float32(mouseX), Y: float32(mouseY)}.ToCell()

	if cellCoord.Valid() {
		cell := g.Board.Cell(cellCoord)
		cell.Focus = true

		if g.LastFocusedCell != nil && g.LastFocusedCell != cell {
			g.LastFocusedCell.Focus = false
		}

		g.LastFocusedCell = cell

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {

			if g.Board.RevealCell(cellCoord, make(map[*Cell]bool)) {
				cell.Exploded = true
				g.Finished = true
				g.Victory = false
			}

		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton2) {
			cell.Flagged = !cell.Flagged

			if cell.Flagged {
				victory := g.Board.CheckVictory()

				if victory {
					g.Finished = true
					g.Victory = true
				}
			}
		}

	} else if g.LastFocusedCell != nil {
		g.LastFocusedCell.Focus = false
		g.LastFocusedCell = nil
	}

	return nil
}

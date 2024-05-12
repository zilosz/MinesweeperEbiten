package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Cell struct {
	HasBomb     bool
	Revealed    bool
	Focus       bool
	Flagged     bool
	Exploded    bool
	BombsNearby int
	Pos         CellCoordinate
}

func (cell *Cell) DrawImage(screen *ebiten.Image, image *ebiten.Image) {
	imageW, imageH := image.Bounds().Dx(), image.Bounds().Dy()

	options := &ebiten.DrawImageOptions{}

	transX := float64(cell.Pos.Col) * float64(imageW)
	transY := float64(cell.Pos.Row) * float64(imageH)
	options.GeoM.Translate(transX, transY)

	scaleX := float64(CellSize) / float64(imageW)
	scaleY := float64(CellSize) / float64(imageH)
	options.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(image, options)
}

func (cell *Cell) Draw(screen *ebiten.Image) {
	screenCoord := cell.Pos.ToScreen()
	x, y := screenCoord.X, screenCoord.Y

	if cell.Revealed {

		if cell.HasBomb {

			if cell.Exploded {
				vector.DrawFilledRect(screen, x, y, CellSize, CellSize, ExplodedBG, false)

			} else {
				vector.DrawFilledRect(screen, x, y, CellSize, CellSize, NormalBG, false)
			}

			cell.DrawImage(screen, BombImage)

		} else {
			vector.DrawFilledRect(screen, x, y, CellSize, CellSize, NormalBG, false)

			if cell.BombsNearby > 0 {
				cell.DrawImage(screen, BombCountImages[cell.BombsNearby])
			}
		}

	} else {
		var bg color.Color

		if cell.Focus {
			bg = FocusBG

		} else {
			bg = HiddenBG
		}

		vector.DrawFilledRect(screen, x, y, CellSize, CellSize, bg, false)

		if cell.Flagged {
			cell.DrawImage(screen, FlagImage)
		}
	}

	vector.StrokeLine(screen, x+CellSize, y, x+CellSize, y+CellSize, CellBorderThickness, color.Black, false)
	vector.StrokeLine(screen, x, y+CellSize, x+CellSize, y+CellSize, CellBorderThickness, color.Black, false)

	if cell.Pos.Row == 0 {
		vector.StrokeLine(screen, x, y, x+CellSize, y, CellBorderThickness, color.Black, false)
	}

	if cell.Pos.Col == 0 {
		vector.StrokeLine(screen, x, y, x, y+CellSize, CellBorderThickness, color.Black, false)
	}
}

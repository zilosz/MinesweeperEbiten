package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	BoardWidth          = 16
	BoardHeight         = 16
	MaxScreenWidth      = 1500
	MaxScreenHeight     = 900
	CellBorderThickness = 4
	BombCount           = 40
)

var (
	HiddenBG   = color.RGBA{R: 156, G: 156, B: 156}
	ExplodedBG = color.RGBA{R: 245, G: 90, B: 90}
	NormalBG   = color.RGBA{R: 195, G: 195, B: 195}
	FocusBG    = color.RGBA{R: 125, G: 125, B: 125}

	BombImage *ebiten.Image
	FlagImage *ebiten.Image

	BombCountImages map[int]*ebiten.Image

	ScreenWidth  int
	ScreenHeight int

	CellSize = float32(MaxScreenWidth) / BoardWidth
)

func LoadImage(path string) *ebiten.Image {
	bombImage, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return bombImage
}

func init() {
	rand.Seed(time.Now().UnixNano())

	BombImage = LoadImage("resources/bomb.png")
	FlagImage = LoadImage("resources/flag.png")

	BombCountImages = make(map[int]*ebiten.Image)

	for i := 1; i < 9; i++ {
		BombCountImages[i] = LoadImage("resources/" + strconv.Itoa(i) + ".png")
	}

	if BoardHeight*CellSize > MaxScreenHeight {
		CellSize = float32(MaxScreenHeight) / BoardHeight
		ScreenWidth, ScreenHeight = int(math.Ceil(float64(CellSize)*BoardWidth)), MaxScreenHeight

	} else {
		ScreenWidth, ScreenHeight = MaxScreenWidth, int(math.Ceil(float64(CellSize)*BoardHeight))
	}
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Minesweeper")

	game := &Game{Board: CreateBoard()}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

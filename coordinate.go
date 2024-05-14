package main

import "math"

type CellCoordinate struct {
	Col int
	Row int
}

type ScreenCoordinate struct {
	X float32
	Y float32
}

func (coordinate CellCoordinate) Valid() bool {
	return coordinate.Col >= 0 && coordinate.Col < BoardWidth && coordinate.Row >= 0 && coordinate.Row < BoardHeight
}

func (coordinate CellCoordinate) ToScreen() ScreenCoordinate {
	return ScreenCoordinate{X: float32(coordinate.Col) * CellSize, Y: float32(coordinate.Row) * CellSize}
}

func (coordinate ScreenCoordinate) ToCell() CellCoordinate {
	column := int(math.Floor(float64(coordinate.X / CellSize)))
	row := int(math.Floor(float64(coordinate.Y / CellSize)))
	return CellCoordinate{Col: column, Row: row}
}

func (coordinate CellCoordinate) Neighbors() []CellCoordinate {
	return []CellCoordinate{
		{Col: coordinate.Col - 1, Row: coordinate.Row - 1},
		{Col: coordinate.Col, Row: coordinate.Row - 1},
		{Col: coordinate.Col + 1, Row: coordinate.Row - 1},
		{Col: coordinate.Col + 1, Row: coordinate.Row},
		{Col: coordinate.Col + 1, Row: coordinate.Row + 1},
		{Col: coordinate.Col, Row: coordinate.Row + 1},
		{Col: coordinate.Col - 1, Row: coordinate.Row + 1},
		{Col: coordinate.Col - 1, Row: coordinate.Row},
	}
}

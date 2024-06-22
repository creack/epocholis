package main

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// Tile holds the list of images to draw in a specific tile position.
type Tile []*Sprite

func (t Tile) Draw(screen *ebiten.Image, g *Game, xi, yi, cx, cy float64, highlight bool) {
	op := &ebiten.DrawImageOptions{}
	for i, s := range t {
		op.GeoM.Reset()
		op.ColorScale.Reset()
		if highlight {
			op.ColorScale.Scale(0, 1, 0, 1)
		}
		op.Blend = ebiten.Blend{}
		op.GeoM.Translate(xi, yi)
		op.GeoM.Translate(s.offsetX, s.offsetY)
		if i == 0 {
			op.Blend = ebiten.BlendDestinationOver
		}
		g.pz.Apply(op)
		op.GeoM.Translate(cx, cy)
		screen.DrawImage(s.Image, op)
	}
}

// Map defines the world by tiles.
type Map [][]Tile

func (l Map) Get(x, y int) Tile {
	if x >= 0 && y >= 0 && x < len(l[0]) && y < len(l) {
		return l[y][x]
	}
	return Tile{}
}

func NewMap(assets *Assets) (Map, error) {
	const w, h = 32, 16
	lines := make(Map, h)
	for y := range lines {
		lines[y] = make([]Tile, w)
		for x, t := range lines[y] {
			t = append(t, assets.Ground)
			switch {
			case x == 1 && y == 1:
				t = append(t, assets.House)

			case x == 0 && y == 0:
				t = append(t, assets.Roads.Get("corner.north.west"))
			case x == 4 && y == 0:
				t = append(t, assets.Roads.Get("cross.east.west.south"))
			case x == 0 && y == 4:
				t = append(t, assets.Roads.Get("cross.north.south.east"))
			case x == 4 && y == 4:
				t = append(t, assets.Roads.Get("cross.full"))
			case y == 0 || y == 4:
				t = append(t, assets.Roads.Get("east.west"))
			case x == 0 || x == 4:
				t = append(t, assets.Roads.Get("north.south"))
			}
			lines[y][x] = t
		}
	}

	return lines, nil
}

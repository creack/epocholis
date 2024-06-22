package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type UnitType string

const (
	UnitTypeEmployeeSeeker = "employee-seeker"
	UnitTypeWaterCarrier   = "water-carrier"
)

type Direction string

const (
	DirectionNorth     = "north"
	DirectionNorthEast = "north.east"
	DirectionEast      = "east"
	DirectionSouthEast = "east.south"
	DirectionSouth     = "south"
	DirectionSouthWest = "south.west"
	DirectionWest      = "west"
)

const (
	spritePerAnimation = 12
	tickPerStep        = 8
)

type UnitAnimation struct {
	g *Game

	direction Direction
	images    Sprites
	count     int
	tmp       int
}

func newUnitAnimation(g *Game, elem Sprites) *UnitAnimation {
	return &UnitAnimation{
		g:      g,
		images: elem,

		count: 0,
		tmp:   0,
	}
}

func (wa *UnitAnimation) Inc() {
	const n = spritePerAnimation * tickPerStep

	wa.count++
	if wa.count%(n+1) == n {
		wa.tmp++
		wa.count = 0
	}
}

func (wa *UnitAnimation) render(target *ebiten.Image, x1, y1, x2, y2 int) {
	const n = spritePerAnimation * tickPerStep
	// Center position.
	cx, cy := float64(wa.g.w/2), float64(wa.g.h/2)

	// Lerp between the source and destination.
	xi1, yi1 := cartesianToIso(tileSize, float64(x1), float64(y1))
	xi2, yi2 := cartesianToIso(tileSize, float64(x2), float64(y2))
	xi, yi := lerp(xi1, yi1, xi2, yi2, float64((wa.count)%(n+1))/n)

	// 5 tick per sprite, 12 sprites per animation.
	i := (wa.count / tickPerStep) % spritePerAnimation

	op := &ebiten.DrawImageOptions{}
	// Move to current isometric position, offset to center the tile.
	op.GeoM.Translate(xi+float64(tileSize)/3, yi-float64(tileSize)/8)
	// Pan/Zoom.
	wa.g.pz.Apply(op)
	// Center.
	op.GeoM.Translate(cx, cy)

	// Render.
	target.DrawImage(wa.images.Filter(wa.direction)[i].Image, op)
}

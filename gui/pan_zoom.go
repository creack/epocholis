package main

import (
	"math"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type panZoom struct {
	mapWidth, mapHeight  int // In tile count.
	camX, camY           float64
	camScale             float64
	camScaleTo           float64
	mousePanX, mousePanY int
}

// Apply the pan/zoom settings to the draw operation.
func (pz *panZoom) Apply(op *ebiten.DrawImageOptions) {
	// Translate camera position.
	op.GeoM.Translate(-pz.camX, pz.camY)
	// Zoom.
	op.GeoM.Scale(pz.camScale, pz.camScale)
}

// Update the settings.
func (pz *panZoom) Update() {
	// Update target zoom level.
	var scrollY float64
	if ebiten.IsKeyPressed(ebiten.KeyC) || ebiten.IsKeyPressed(ebiten.KeyPageDown) {
		scrollY = -0.25
	} else if ebiten.IsKeyPressed(ebiten.KeyE) || ebiten.IsKeyPressed(ebiten.KeyPageUp) {
		scrollY = .25
	} else {
		_, scrollY = ebiten.Wheel()
		if scrollY < -1 {
			scrollY = -1
		} else if scrollY > 1 {
			scrollY = 1
		}
	}
	pz.camScaleTo += scrollY * (pz.camScaleTo / 7)

	// Clamp target zoom level.
	if pz.camScaleTo < 0.01 {
		pz.camScaleTo = 0.01
	} else if pz.camScaleTo > 100 {
		pz.camScaleTo = 100
	}

	// Smooth zoom transition.
	const div = 30.0
	if pz.camScaleTo > pz.camScale {
		pz.camScale += (pz.camScaleTo - pz.camScale) / div
	} else if pz.camScaleTo < pz.camScale {
		pz.camScale -= (pz.camScale - pz.camScaleTo) / div
	}

	// Pan camera via keyboard.
	pan := 7.0 / pz.camScale
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		pz.camX -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		pz.camX += pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		pz.camY -= pan
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		pz.camY += pan
	}

	// Pan camera via mouse.
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if pz.mousePanX == math.MinInt32 && pz.mousePanY == math.MinInt32 {
			pz.mousePanX, pz.mousePanY = ebiten.CursorPosition()
		} else {
			x, y := ebiten.CursorPosition()
			dx, dy := float64(pz.mousePanX-x)*(pan/100), float64(pz.mousePanY-y)*(pan/100)
			pz.camX, pz.camY = pz.camX-dx, pz.camY+dy
		}
	} else if pz.mousePanX != math.MinInt32 || pz.mousePanY != math.MinInt32 {
		pz.mousePanX, pz.mousePanY = math.MinInt32, math.MinInt32
	}

	// Clamp camera position.
	worldWidth := float64(pz.mapWidth * tileSize / 2)
	if pz.camX < -worldWidth {
		pz.camX = -worldWidth
	} else if pz.camX > worldWidth {
		pz.camX = worldWidth
	}
	worldHeight := float64(pz.mapHeight * tileSize / 2)
	if pz.camY < -worldHeight {
		pz.camY = -worldHeight
	} else if pz.camY > 0 {
		pz.camY = 0
	}
}

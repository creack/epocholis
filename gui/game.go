package main

import (
	"fmt"
	"math"
	"strings"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game is an isometric demo game.
type Game struct {
	assets           *Assets                     // Pre-loaded assets.
	workerAnimations map[UnitType]*UnitAnimation // Pre-loaded animations.

	gameMap Map // Game map state.

	w, h int     // Width/Height of the window.
	pz   panZoom // Pan/Zoom settings.

	cursorIX, cursorIY float64
}

// NewGame returns a new isometric demo Game.
func NewGame(assets *Assets) (*Game, error) {
	m, err := NewMap(assets)
	if err != nil {
		return nil, fmt.Errorf("newMap: %w", err)
	}
	if len(m) == 0 {
		return nil, fmt.Errorf("missing map")
	}
	return &Game{
		assets:  assets,
		gameMap: m,

		pz: panZoom{
			mapWidth:   len(m[0]),
			mapHeight:  len(m),
			mousePanX:  math.MinInt32,
			mousePanY:  math.MinInt32,
			camScale:   2.81,
			camScaleTo: 2.81,
			camX:       66,
			camY:       -66,
			// camScale:   0.01,
			// camScaleTo: 0.94,
			// camX:       250,
			// camY:       -360,
		},
	}, nil
}

// Update reads current user input and updates the Game state.
func (g *Game) Update() error {
	for _, elem := range g.workerAnimations {
		elem.Inc()
	}
	dx, dy := ebiten.CursorPosition()
	cx, cy := float64(g.w/2), float64(g.h/2)
	// Apply the geom in reverse.
	ix, iy := isoToCartesian(tileSize, (float64(dx)-cx)/g.pz.camScale+g.pz.camX, (float64(dy)-cy)/g.pz.camScale-g.pz.camY)
	g.cursorIX = math.Round(ix) - 1
	g.cursorIY = math.Round(iy)

	g.pz.Update()
	return nil
}

// Draw draws the Game on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Render the world tiles.
	g.renderTiles(screen)
	// Render the animation sprites.
	g.renderAnimations(screen)

	// Print game info.
	msgs := []string{
		"KEYS WASD EC R",
		fmt.Sprintf("FPS  %0.0f", ebiten.ActualFPS()),
		fmt.Sprintf("TPS  %0.0f", ebiten.ActualTPS()),
		fmt.Sprintf("SCA  %0.2f", g.pz.camScale),
		fmt.Sprintf("POS  %0.0f,%0.0f", g.pz.camX, g.pz.camY),
		"",
		fmt.Sprintf("CUR %.3f,%.3f", g.cursorIX, g.cursorIY),
	}
	ebitenutil.DebugPrint(screen, strings.Join(msgs, "\n"))
}

// Layout is called when the Game's layout changes.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.w, g.h = outsideWidth, outsideHeight
	return g.w, g.h
}

func (g *Game) renderTiles(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	cx, cy := float64(g.w/2), float64(g.h/2)
	_, _ = cx, cy
	for y, line := range g.gameMap {
		for x, t := range line {
			xi, yi := cartesianToIso(tileSize, float64(x), float64(y))

			// op.GeoM.Reset()
			// op.ColorScale.Reset()
			// // Move to current isometric position.
			// op.GeoM.Translate(xi, yi)
			// op.GeoM.Translate(t.offsetX, t.offsetY)
			// // Pan/Zoom.
			// g.pz.Apply(op)
			// // Center.
			// op.GeoM.Translate(cx, cy)

			// Highlight cursor.
			if x == int(g.cursorIX) && y == int(g.cursorIY) {
				op.ColorScale.Scale(0, 1, 0, 1)
			}

			t.Draw(screen, g, xi, yi, cx, cy, x == int(g.cursorIX) && y == int(g.cursorIY))

		}
	}
}

func (g *Game) renderAnimations(screen *ebiten.Image) {
	wa := g.workerAnimations[UnitTypeWaterCarrier]
	{
		wa.direction = DirectionNorth
		wa.tmp = 0
		wa.render(screen, 0, 5-wa.tmp%5, 0, 5-wa.tmp%5)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\n\n\n\n\n\nTMP: %d\nCNT: %d\nI: %d\n", wa.tmp, wa.count, (wa.count/tickPerStep)%spritePerAnimation))
	return
	{
		wa.direction = DirectionSouth
		wa.render(screen, 0, 0+wa.tmp%5, 0, 1+wa.tmp%5)
	}
	{
		wa.direction = DirectionEast
		wa.render(screen, 1+wa.tmp%5, 0, 2+wa.tmp%5, 0)
	}
	{
		wa.direction = DirectionWest
		wa.render(screen, 5-wa.tmp%5, 0, 4-wa.tmp%5, 0)
	}
}

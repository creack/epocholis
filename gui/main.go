package main

import (
	"log"
	"os"
	"strconv"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// dx, dy := cartesianToIso(tileSize, 1, 1)
	// fmt.Println(dx, dy)
	// fmt.Println(isoToCartesian(tileSize, dx, dy))
	// return
	ebiten.SetWindowTitle("Epocholis")
	ebiten.SetWindowSize(1280, 700)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	//	ebiten.SetTPS(10)

	// Open the window in the first monitor. (Useful when working with multi monitor).
	if monitorStr := os.Getenv("MONITOR"); monitorStr != "" {
		monitor, _ := strconv.Atoi(monitorStr)
		monitors := ebiten.AppendMonitors(nil)
		ebiten.SetMonitor(monitors[monitor])
		ebiten.SetWindowPosition(0, 0)
	}

	// Load the embeded assets into ebiten.Image's.
	assets, err := LoadAssets()
	if err != nil {
		log.Fatalf("LoadAssets: %s.", err)
	}

	g, err := NewGame(assets)
	if err != nil {
		log.Fatalf("NewGame: %s", err)
	}

	g.workerAnimations = map[UnitType]*UnitAnimation{}
	for unitType, sprites := range g.assets.Units {
		g.workerAnimations[unitType] = newUnitAnimation(g, sprites)
	}
	// for dir, elem := range assets.Units[UnitTypeEmployeeSeeker] {
	// 	g.workerAnimations[string(dir)] = newWorkerAnimation(g, elem)
	// }

	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("RunGame: %s.", err)
	}
}

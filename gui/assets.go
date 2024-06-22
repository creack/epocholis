package main

import (
	"embed"
	"fmt"
	_ "image/png" // Load the png driver.
	"path"
)

// Asset tiles are 60x60 ish.
const tileSize = 60

//go:embed assets
var rawAssets embed.FS

// Assets holds all the loaded assets.
type Assets struct {
	Units map[UnitType]Sprites

	Roads  Sprites
	Ground *Sprite
	House  *Sprite
}

func LoadAssets() (*Assets, error) {
	assets := &Assets{}

	// Load the roads.
	roadSprites, err := NewSprites(rawAssets, "assets/roads")
	if err != nil {
		return nil, fmt.Errorf("newSprites assets/roads: %w", err)
	}
	assets.Roads = roadSprites

	// Load the ground.
	{
		sprite, err := NewSprite(rawAssets, "assets/ground.png")
		if err != nil {
			return nil, fmt.Errorf("newSprite ground.png: %w", err)
		}
		assets.Ground = sprite
	}
	// Load Houses.
	{
		sprite, err := NewSprite(rawAssets, "assets/house.3-1.png")
		if err != nil {
			return nil, fmt.Errorf("newSprite house.1-1.png: %w", err)
		}
		sprite.offsetX = -tileSize / 2
		sprite.offsetY = -tileSize / 2
		// sprite.offsetX = 8
		// sprite.offsetY = -14
		assets.House = sprite
	}

	// Load units.
	entries, err := rawAssets.ReadDir("assets/units")
	if err != nil {
		return nil, fmt.Errorf("readDir assets/units: %w", err)
	}
	assets.Units = map[UnitType]Sprites{}
	for _, entry := range entries {
		unitType := UnitType(entry.Name())

		dirPath := path.Join("assets", "units", entry.Name())
		sprites, err := NewSprites(rawAssets, dirPath)
		if err != nil {
			return nil, fmt.Errorf("newSprites %q: %w", dirPath, err)
		}
		assets.Units[unitType] = sprites
	}

	return assets, nil
}

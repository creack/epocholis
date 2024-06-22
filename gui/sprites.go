package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"path"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite wraps the imaeg with it's asset name.
type Sprite struct {
	*ebiten.Image
	name             string
	offsetX, offsetY float64
}

// NewSprite loads the given sprite assets into a Sprite object.
func NewSprite(raw embed.FS, filePath string) (*Sprite, error) {
	buf, err := raw.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("readFile: %w", err)
	}
	img, _, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}
	return &Sprite{
		Image: ebiten.NewImageFromImage(img),
		name:  strings.TrimSuffix(path.Base(filePath), path.Ext(filePath)),
	}, nil
}

// List of sprites. Not using a map as we want to maintain order for animations.
type Sprites []*Sprite

// NewSprites loads all the sprites assets in the given path.
func NewSprites(raw embed.FS, basePath string) (Sprites, error) {
	var out Sprites
	entries, err := rawAssets.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("readDir: %w", err)
	}
	for _, elem := range entries {
		sprite, err := NewSprite(rawAssets, path.Join(basePath, elem.Name()))
		if err != nil {
			return nil, fmt.Errorf("newSprite %q: %w", elem.Name(), err)
		}
		out = append(out, sprite)
	}
	return out, nil
}

func (s Sprites) Get(name string) *Sprite {
	for _, elem := range s {
		if elem.name == name {
			return elem
		}
	}
	panic("not found: " + name)
}

func (s Sprites) Filter(dir Direction) Sprites {
	var out Sprites
	for _, elem := range s {
		if strings.HasPrefix(elem.name, string(dir)) {
			out = append(out, elem)
		}
	}
	return out
}

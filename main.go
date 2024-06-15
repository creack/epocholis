package main

import (
	"fmt"
	"time"
)

func clearScreen() {
	fmt.Printf("\033c")
}

type gridCase interface {
	render() string
}

type grid [][]gridCase

type empty struct{}

func (empty) render() string { return " " }

type controller struct {
	grid grid
}

func newController(width, height int) *controller {
	c := &controller{}
	c.grid = make(grid, height)
	for i := range c.grid {
		c.grid[i] = make([]gridCase, width)
		for j := range c.grid[i] {
			c.grid[i][j] = empty{}
		}
	}
	return c
}

func (c *controller) render() {
	var buf string
	buf += "┌"
	for range c.grid[0] {
		buf += "─"
	}
	buf += "┐\n"
	for _, line := range c.grid {
		buf += "│"
		for _, elem := range line {
			buf += elem.render()
		}
		buf += "│\n"
	}
	buf += "└"
	for range c.grid[0] {
		buf += "─"
	}
	buf += "┘\n"
	clearScreen()
	fmt.Printf("%s", buf)
}

func run() error {
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	c := newController(32, 16)

	for i := range 16 {
		newRoad(c.newBase(6+i, 2))
	}

	h := newHouse(c.newBase(10, 3))
	h2 := newHouse(c.newBase(16, 3))
loop:
	c.render()
	h.tick()
	h2.tick()
	<-ticker.C
	goto loop
}

func main() {
	if err := run(); err != nil {
		println("Fail:", err.Error())
		return
	}
}

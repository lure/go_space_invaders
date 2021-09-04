package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL:", err)
		return
	}

	// ------------------------------------------------------------------------------------------------
	// WINDOWS
	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Creating window:", err)
		return
	}
	window.Show()

	defer window.Destroy()

	// ------------------------------------------------------------------------------------------------
	// RENDERER
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing render:", err)
		return
	}

	defer renderer.Destroy()

	// ------------------------------------------------------------------------------------------------
	// PLAYER
	elements = append(elements, newPlayer(renderer))

	// ------------------------------------------------------------------------------------------------
	// ENEMIES
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i) / 5) * screenWidth + (beSize / 2)
			y := float64(j) * beSize + (beSize / 2)
			enemy := newBasicEnemy(renderer, vector { x: x, y: y})
			elements = append(elements, enemy)
		}
	}

	// ------------------------------------------------------------------------------------------------
	// BULLETS
	initBulletPool(renderer)

	// ------------------------------------------------------------------------------------------------
	// GAME LOOP
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, elem := range elements {
			if elem.active {
				elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}

				elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
					return
				}
			}
		}

		renderer.Present()
	}
}


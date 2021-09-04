package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	playerSpeed        = 0.05
	playerShotCooldown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) *element {
	e := &element{}

	sr := newSpriteRenderer(e, renderer, "sprites/player.bmp")
	e.addComponent(sr)

	e.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - sr.height/2.0,
	}
	e.active = true


	mover := newKeyboardMover(e, playerSpeed)
	e.addComponent(mover)

	shooter := newKeyboardShooter(e, playerShotCooldown)
	e.addComponent(shooter)

	return e
}

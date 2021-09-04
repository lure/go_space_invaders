package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	beSize = 105
)

func newBasicEnemy(renderer *sdl.Renderer, v vector) *element {
	e := &element{}
	e.position = v
	e.angle = 0 //180

	e.active = true
	sr := newSpriteRenderer(e, renderer, "sprites/invader_b.bmp")
	e.addComponent(sr)

	return e
}

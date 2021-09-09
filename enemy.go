package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	beSize = 105
)

func newBasicEnemy(renderer *sdl.Renderer, v vector) *element {
	e := &element{}
	e.position = v
	e.angle = 0 //180

	col := circle{
		center: e.position,
		radius: 38,
	}
	e.collisions = append(e.collisions, col)

	e.active = true
	sr := newSpriteRenderer(e, renderer, "sprites/invader_b.bmp")
	e.addComponent(sr)

	vtp := newVulnerableToProjectiles(e)
	e.addComponent(vtp)

	return e
}

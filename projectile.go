package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (sr *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (sr *bulletMover) onUpdate() error {
	bul := sr.container

	bul.position.x += bulletSpeed * math.Cos(bul.angle)
	bul.position.y += bulletSpeed * math.Sin(bul.angle)

	if bul.position.x > screenWidth || bul.position.x < 0 ||
		bul.position.y < 0 || bul.position.y > screenHeight {
		bul.active = false
	}

	return nil
}

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

func (c *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (c *bulletMover) onUpdate() error {
	bul := c.container

	bul.position.x += bulletSpeed * math.Cos(bul.angle)
	bul.position.y += bulletSpeed * math.Sin(bul.angle)

	if bul.position.x > screenWidth || bul.position.x < 0 ||
		bul.position.y > screenHeight || bul.position.y < 0 {
		bul.active = false
	}

	bul.collisions[0].center = bul.position

	return nil
}

func (c *bulletMover) onCollision(other *element) error {
	if other.tag != "bullet" {
		c.container.active = false
	}
	return nil
}

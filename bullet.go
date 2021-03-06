package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSpeed = 15
)

func newBullet(renderer *sdl.Renderer) *element {
	e := &element{}
	e.tag = "bullet"
	sr := newSpriteRenderer(e, renderer, "sprites/bullet.bmp")
	e.addComponent(sr)
	mover := newBulletMover(e, bulletSpeed)
	e.addComponent(mover)

	col := circle{
		center: e.position,
		radius: 8,
	}
	e.collisions = append(e.collisions, col)
	return e
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullet := newBullet(renderer)
		bulletPool = append(bulletPool, bullet)
		elements = append(elements, bullet)
	}
}

// bool here is a signal of failed invocation which is not as critical as error
func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}

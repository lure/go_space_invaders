package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"time"
)

type keyboardMover struct {
	container *element
	speed     float64

	sr *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {

	keys := sdl.GetKeyboardState()
	cont := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.position.x-mover.sr.width/2 > 0 {
			cont.position.x -= mover.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.position.x+mover.sr.width/2 < screenWidth {
			cont.position.x += mover.speed
		}
	}
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			pos := shooter.container.position
			shooter.shoot(pos.x+25, pos.y-20)
			shooter.shoot(pos.x-25, pos.y-20)
			shooter.lastShot = time.Now()
		}
	}
	return nil
}

func (shooter *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) onCollision(other *element) error {
	return nil
}

func (shooter *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.position.x = x
		bul.position.y = y
		bul.angle = 270 * (math.Pi / 180)
	}
}

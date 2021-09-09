package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type vulnerableToProjectiles struct {
	container *element
}

func newVulnerableToProjectiles(container *element) *vulnerableToProjectiles {
	return &vulnerableToProjectiles{
		container: container,
	}
}

func (e *vulnerableToProjectiles) onUpdate() error {
	return nil
}

func (e *vulnerableToProjectiles) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (e *vulnerableToProjectiles) onCollision(other *element) error {
	if other.tag == "bullet" {
		e.container.active = false
	}

	return nil
}

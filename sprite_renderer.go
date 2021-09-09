package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	tex           *sdl.Texture
	width, height float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	tex := textureFromBMP(renderer, filename)
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("quering texture %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       tex,
		width:     float64(width),
		height:    float64(height),
	}
}

func (c *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	x := c.container.position.x - float64(c.width)/2.0
	y := c.container.position.y - float64(c.height)/2.0

	err := renderer.CopyEx(
		c.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(c.width), H: int32(c.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(c.width), H: int32(c.height)},
		c.container.angle,
		&sdl.Point{X: int32(c.width / 2), Y: int32(c.height / 2)},
		sdl.FLIP_NONE,
	)

	return err
}

func (c *spriteRenderer) onUpdate() error {
	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}

	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))

	}
	return texture
}

func (c *spriteRenderer) onCollision(other *element) error {
	return nil
}

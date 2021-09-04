package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type vector struct {
	x, y float64
}

type element struct {
	position   vector
	angle      float64
	active     bool
	components []component
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (elem *element) addComponent(c component) {
	t := reflect.TypeOf(c)
	for _, existing := range elem.components {
		if reflect.TypeOf(existing) == t {
			panic(fmt.Errorf("attempt to add new component with exsisting type %v", t))
		}
	}
	elem.components = append(elem.components, c)
}

func (elem *element) getComponent(withType component) component {
	t := reflect.TypeOf(withType)
	for _, existing := range elem.components {
		if reflect.TypeOf(existing) == t {
			return existing
		}
	}
	panic(fmt.Errorf("no component with exsisting type %v", t))
}

// contains all drawable, controllable, updatable things in the game
var elements []*element

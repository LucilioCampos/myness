package ui

import (
	"image"

	"github.com/LucilioCampos/mynes/mynes"
	"github.com/go-gl/gl/v2.1/gl"
)

type GameView struct {
	director *Director
	console  *mynes.Console
	title    string
	hash     string
	texture  uint32
	record   bool
	frames   []image.Image
}

// Enter implements View.
func (view *GameView) Enter() {

}

// Exit implements View.
func (view *GameView) Exit() {

}

// Update implements View.
func (view *GameView) Update(t float64, dt float64) {
	if dt > 1 {
		dt = 0
	}
	console := view.console
	console.StepSeconds(dt)
	gl.BindTexture(gl.TEXTURE_2D, view.texture)
}

func NewGameView(director *Director, console *mynes.Console, title, hash string) View {
	texture := createTexture()
	return &GameView{director, console, title, hash, texture, false, nil}
}

func createTexture() uint32 {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return texture
}

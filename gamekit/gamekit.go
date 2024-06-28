package gamekit

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"image/png"
)

type DefaultLayout struct{}

func (l DefaultLayout) WindowSize() (int, int) {
	return 1920 / 2, 1080 / 2
}

func (l DefaultLayout) Layout(int, int) (int, int) {
	return l.WindowSize()
}

func LoadShader(filename string) *ebiten.Shader {
	src, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	shader, err := ebiten.NewShader(src)
	if err != nil {
		panic(err)
	}
	return shader
}

func LoadImage(filename string) *ebiten.Image {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pngImage, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(pngImage)
}

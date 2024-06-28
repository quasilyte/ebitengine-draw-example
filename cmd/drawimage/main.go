package main

import (
	"mygame/gamekit"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &myGame{}

	ebiten.SetWindowTitle("DrawImage example")
	ebiten.SetWindowSize(g.WindowSize())
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type myGame struct {
	initialized bool
	gamekit.DefaultLayout
	objects []*object
}

type object struct {
	image *ebiten.Image
	pos   [2]float64
}

func (g *myGame) Update() error {
	if !g.initialized {
		g.initialized = true
		g.init()
	}
	return nil
}

func (g *myGame) init() {
	img := gamekit.LoadImage("gopher.png")
	for i := 0; i < 9; i++ {
		o := &object{
			image: img,
			pos:   [2]float64{float64(i*96) + 64, 128},
		}
		g.objects = append(g.objects, o)
	}
}

func (g *myGame) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		var opts ebiten.DrawImageOptions
		opts.GeoM.Translate(o.pos[0], o.pos[1])
		screen.DrawImage(o.image, &opts)
	}
}

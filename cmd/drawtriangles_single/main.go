package main

import (
	"mygame/gamekit"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &myGame{}

	ebiten.SetWindowTitle("draw triangles example")
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
	pos   [2]float32
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
			pos:   [2]float32{float32(i*96) + 64, 128},
		}
		g.objects = append(g.objects, o)
	}
}

func (g *myGame) Draw(screen *ebiten.Image) {
	for _, o := range g.objects {
		img := o.image
		iw, ih := img.Size()
		w := float32(iw)
		h := float32(ih)

		// Здесь было бы много преобразований float64->float32,
		// но я отредактировал object так, чтобы уместить
		// код по ширине для комфорта хабравчан.
		vertices := []ebiten.Vertex{
			{
				DstX: o.pos[0], DstY: o.pos[1],
				SrcX: 0, SrcY: 0,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			{
				DstX: o.pos[0] + w, DstY: o.pos[1],
				SrcX: float32(w), SrcY: 0,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			{
				DstX: o.pos[0], DstY: o.pos[1] + h,
				SrcX: 0, SrcY: h,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			{
				DstX: o.pos[0] + w, DstY: o.pos[1] + h,
				SrcX: w, SrcY: h,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
		}

		indices := []uint16{
			0, 1, 2,
			1, 2, 3,
		}

		var opts ebiten.DrawTrianglesOptions
		screen.DrawTriangles(vertices, indices, img, &opts)
	}
}

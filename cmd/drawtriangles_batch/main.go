package main

import (
	"mygame/gamekit"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &myGame{}

	ebiten.SetWindowTitle("draw triangles batch example")
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
	// Здесь мы пользуемся фактом, что все объекты у нас используют
	// одно и то же изображение.

	img := g.objects[0].image
	iw, ih := img.Size()
	w := float32(iw)
	h := float32(ih)

	// Аллоцируем вершинки и индексы сразу для всех объектов.
	vertices := make([]ebiten.Vertex, 0, 4*len(g.objects))
	indices := make([]uint16, 0, 6*len(g.objects))

	i := uint16(0)

	for _, o := range g.objects {
		vertices = append(vertices,
			ebiten.Vertex{
				DstX: o.pos[0], DstY: o.pos[1],
				SrcX: 0, SrcY: 0,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			ebiten.Vertex{
				DstX: o.pos[0] + w, DstY: o.pos[1],
				SrcX: float32(w), SrcY: 0,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			ebiten.Vertex{
				DstX: o.pos[0], DstY: o.pos[1] + h,
				SrcX: 0, SrcY: h,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
			ebiten.Vertex{
				DstX: o.pos[0] + w, DstY: o.pos[1] + h,
				SrcX: w, SrcY: h,
				ColorR: 1, ColorG: 1,
				ColorB: 1, ColorA: 1,
			},
		)

		indices = append(indices,
			i+0, i+1, i+2,
			i+1, i+2, i+3,
		)
		i += 4 // Увеличиваем на количество vertices на объект
	}

	var opts ebiten.DrawTrianglesOptions
	screen.DrawTriangles(vertices, indices, img, &opts)
}

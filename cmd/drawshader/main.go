package main

import (
	"mygame/gamekit"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &myGame{}

	ebiten.SetWindowTitle("draw shader example")
	ebiten.SetWindowSize(g.WindowSize())
	ebiten.SetFullscreen(true)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

type myGame struct {
	initialized bool
	gamekit.DefaultLayout
	circles []*circle
	rects   []*rectangle
}

type rectangle struct {
	shader *ebiten.Shader
	width  float32
	height float32
	pos    [2]float32
}

type circle struct {
	shader   *ebiten.Shader
	uniforms map[string]any
	radius   float32
	pos      [2]float32
}

func (g *myGame) Update() error {
	if !g.initialized {
		g.initialized = true
		g.init()
	}
	return nil
}

func (g *myGame) init() {
	circleShader := gamekit.LoadShader("circle_shader.go")
	rectShader := gamekit.LoadShader("rect_shader.go")

	for i := 0; i < 9; i++ {
		r := &rectangle{
			shader: rectShader,
			width:  40,
			height: 40,
			pos:    [2]float32{float32(i*96) + 64, 128},
		}
		g.rects = append(g.rects, r)

		c := &circle{
			pos:      [2]float32{float32(i*96) + 64, 128},
			shader:   circleShader,
			radius:   16,
			uniforms: map[string]any{"Radius": float32(16)},
		}
		g.circles = append(g.circles, c)
	}
}

func (g *myGame) Draw(screen *ebiten.Image) {
	for _, r := range g.rects {
		g.drawShader(screen, drawShaderOptions{
			shader: r.shader,
			pos:    r.pos,
			width:  r.width,
			height: r.height,
		})
	}

	for _, c := range g.circles {
		g.drawShader(screen, drawShaderOptions{
			shader:   c.shader,
			uniforms: c.uniforms,
			pos:      c.pos,
			width:    2 * c.radius,
			height:   2 * c.radius,
		})
	}
}

type drawShaderOptions struct {
	pos      [2]float32
	shader   *ebiten.Shader
	width    float32
	height   float32
	uniforms map[string]any
}

func (g *myGame) drawShader(dst *ebiten.Image, opts drawShaderOptions) {
	pos := opts.pos
	w := opts.width
	h := opts.height

	// Будем рисовать относительно центра.
	pos[0] -= w / 2
	pos[1] -= h / 2

	vertices := []ebiten.Vertex{
		{
			DstX: pos[0], DstY: pos[1],
			SrcX: 0, SrcY: 0,
			ColorR: 1, ColorG: 1,
			ColorB: 1, ColorA: 1,
		},
		{
			DstX: pos[0] + w, DstY: pos[1],
			SrcX: float32(w), SrcY: 0,
			ColorR: 1, ColorG: 1,
			ColorB: 1, ColorA: 1,
		},
		{
			DstX: pos[0], DstY: pos[1] + h,
			SrcX: 0, SrcY: h,
			ColorR: 1, ColorG: 1,
			ColorB: 1, ColorA: 1,
		},
		{
			DstX: pos[0] + w, DstY: pos[1] + h,
			SrcX: w, SrcY: h,
			ColorR: 1, ColorG: 1,
			ColorB: 1, ColorA: 1,
		},
	}

	indices := []uint16{
		0, 1, 2,
		1, 2, 3,
	}

	var drawOptions ebiten.DrawTrianglesShaderOptions
	drawOptions.Uniforms = opts.uniforms
	dst.DrawTrianglesShader(vertices, indices, opts.shader, &drawOptions)
}

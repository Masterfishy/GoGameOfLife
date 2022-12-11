package main

import (
	"time"
	"math/rand"

	. "github.com/Masterfishy/GopherLife/engine"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var square = []float32{
    -1, 1, 0,
    -1, -1, 0,
    1, -1, 0,

    -1, 1, 0,
    1, 1, 0,
    1, -1, 0,
}

type GopherLifeGame struct {
	engine *Engine
	window *glfw.Window
	Rows int
	Cols int
	Seed int64
	Threshold float64
	Fps int
	RenderSystem *RenderSystem
	LivingSystem *LivingSystem
}

func NewGopherLifeGame(engine *Engine, window *glfw.Window, rows int, cols int, seed int64, threshold float64, fps int) (*GopherLifeGame, error) {
	g := &GopherLifeGame{
		engine: engine,
		window: window,
		Rows: rows,
		Cols: cols,
		Seed: seed,
		Threshold: threshold,
		Fps: fps,
	}

	return g, nil
}

// Start sets up the the GopherLifeGame
func (g GopherLifeGame) SetUp() {
	// Create the systems
	g.RenderSystem, _ = NewRenderSystem(g.window)
	g.LivingSystem, _ = NewLivingSystem(g.Rows, g.Cols)

	g.engine.RenderSystem = g.RenderSystem
	g.engine.LivingSystem = g.LivingSystem
}

// Play begins the the GopherLifeGame
func (g GopherLifeGame) Play() {
    g.entityCreator()

	for !g.window.ShouldClose() {
        t := time.Now()

        g.engine.Update(0)

        time.Sleep(time.Second/time.Duration(g.Fps) - time.Since(t))
    }
}

func (g GopherLifeGame) entityCreator() {
	rand.Seed(g.Seed)

    for x := 0; x < g.Rows; x++ {
        for y := 0; y < g.Cols; y++ {
            c := newEntity(x, y)

            g.engine.AddEntity(c)
        }
    }
}

func newEntity(x, y int) *Entity {
    points := make([]float32, len(square))
    copy(points, square)

    for i := 0; i < len(points); i++ {
        var position float32
        var size float32
        
        switch i % 3 {
        case 0:
            size = 1.0 / float32(cols)
            position = float32(x) * size
        case 1:
            size = 1.0 / float32(rows)
            position = float32(y) * size
        default:
            continue
        }

        if points[i] < 0 {
            points[i] = (position * 2) - 1
        } else {
            points[i] = ((position + size) * 2) - 1
        }
    }

    alive := rand.Float64() < threshold

    // Create components
    positionComponent := &PositionComponent{ X: float32(x), Y: float32(y), Rotation: 0 }
    displayComponent := &DisplayComponent{ Points: points, X: float32(x), Y: float32(y), Rotation: 0 }
    livingComponent := &LivingComponent{ Alive: alive, AliveNext: alive }

    // Create entity
    var cell Entity
    cell.Position = positionComponent
    cell.Display = displayComponent
    cell.Living = livingComponent

	return &cell
}
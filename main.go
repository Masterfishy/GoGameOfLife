package main

import (
	"runtime"
    "math/rand"
    "time"
    "flag"

    . "papaya/engine"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
    width  = 500
    height = 500
)

var (
    square = []float32{
        -0.5, 0.5, 0,
        -0.5, -0.5, 0,
        0.5, -0.5, 0,
    
        -0.5, 0.5, 0,
        0.5, 0.5, 0,
        0.5, -0.5, 0,
    }

    rows = 20
    cols = 20
    seed = time.Now().UnixNano()
    threshold = 0.2
    fps = 10

    renderSystem = new(RenderSystem)
    livingSystem = new(LivingSystem)
)

func init() {
    flag.IntVar(&cols, "columns", cols, "Sets the number of columns.")
	flag.IntVar(&rows, "rows", rows, "Sets the number of rows.")
	flag.Int64Var(&seed, "seed", seed, "Sets the starting seed of the game, used to randomize the initial state.")
	flag.Float64Var(&threshold, "threshold", threshold, "A percentage between 0 and 1 used in conjunction with the -seed to determine if a cell starts alive. For example, 0.15 means each cell has a 15% chance of starting alive.")
	flag.IntVar(&fps, "fps", fps, "Sets the frames-per-second, used set the speed of the simulation.")
	flag.Parse()

    // This is needed to arrange that main() runs on main thread.
    // See documentation for functions that are only allowed to be called from the main thread.
    runtime.LockOSThread()
}

func main() {
    // Initialize GLFW and OpenGL
    window := initGlfw()
    defer glfw.Terminate()
    
    renderSystem.Window = window
    renderSystem.Start()

    livingSystem.Targets = make([][]LivingNode, rows, cols)

    makeCells()

    for !window.ShouldClose() {
        t := time.Now()

        livingSystem.Update(0)
        renderSystem.Update(0)

        time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
    }
}

// Create cells in rows and cols
func makeCells() [][]*Entity {
    rand.Seed(seed)

    cells := make([][]*Entity, rows, cols)

    for x := 0; x < rows; x++ {
        for y := 0; y < cols; y++ {
            c := newCell(x, y)
            cells[x] = append(cells[x], c)
        }
    }

    return cells
}

func newCell(x, y int) *Entity {
    points := make([]float32, len(square), len(square))
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

    positionComponent := &PositionComponent{ X: float32(x), Y: float32(y), Rotation: 0 }
    displayComponent := &DisplayComponent{ Points: points, X: float32(x), Y: float32(y), Rotation: 0 }
    livingComponent := &LivingComponent{ Alive: alive, AliveNext: alive }

    renderNode := &RenderNode{ 
        Position: positionComponent,
        Display: displayComponent,
        Living: livingComponent,
    }

    renderSystem.Targets = append(renderSystem.Targets, *renderNode)

    livingNode := &LivingNode{
        Living: livingComponent,
        Position: positionComponent,
    }

    livingSystem.Targets[x] = append(livingSystem.Targets[x], *livingNode)

    var cell Entity
    cell.Components = make(map[string]any)
    cell.AddComponent("position", positionComponent)
    cell.AddComponent("display", displayComponent)
    cell.AddComponent("living", livingComponent)

    return &cell
}

// Initialize glfw and return a window to use.
func initGlfw() *glfw.Window {
    if err := glfw.Init(); err != nil {
        panic(err)
    }

    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(width, height, "Testing", nil, nil)
    if err != nil {
        panic(err)
    }

    window.MakeContextCurrent()

    return window
}

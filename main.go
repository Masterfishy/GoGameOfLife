package main

import (
	"runtime"
    "time"
    "flag"

    "github.com/Masterfishy/GopherLife/engine"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
    width  = 500
    height = 500
)

var (
    rows = 20
    cols = 20
    seed = time.Now().UnixNano()
    threshold = 0.2
    fps = 10

    gameEngine *engine.Engine
    game *GopherLifeGame
)

func init() {
    flag.IntVar(&cols, "cols", cols, "Sets the number of columns.")
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

    // Create the engine
    gameEngine, _ = engine.NewEngine()

    // Create the game and play
    game, _ = NewGopherLifeGame(gameEngine, window, rows, cols, seed, threshold, fps)
    game.SetUp()
    game.Play()
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

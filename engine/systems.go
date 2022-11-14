package engine

import (
	. "papaya/graphics"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type ISystem interface {
	Start()
	Update(time float32)
	End()
}

type MoveSystem struct {
	Targets []MoveNode
}

func (ms MoveSystem) Update(time float32) {
	for _, target := range ms.Targets {
		target.Position.X += target.Velocity.X * time;
		target.Position.Y += target.Velocity.Y * time;
		target.Position.Rotation += target.Velocity.AngularVelocity * time;
	}
}

///////////////////
// Render System //
///////////////////
type RenderSystem struct {
	Targets []RenderNode
	Window *glfw.Window
	Program uint32
}

// Initialize OpenGL
func (rs RenderSystem) Start() {
	// Initialize OpenGL
	if err := gl.Init(); err != nil {
        panic(err)
    }

    vertexShader, err := CompileShader(VertexShaderSource, gl.VERTEX_SHADER)
    if err != nil {
        panic(err)
    }

    fragmentShader, err := CompileShader(FragmentShaderSource, gl.FRAGMENT_SHADER)
    if err != nil {
        panic(err)
    }

    program := gl.CreateProgram()
    gl.AttachShader(program, vertexShader)
    gl.AttachShader(program, fragmentShader)
    gl.LinkProgram(program)

    rs.Program = program
}

// Draw all Render Nodes
func (rs RenderSystem) Update(time float32) {
	// Clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.UseProgram(rs.Program)

	for _, target := range rs.Targets {
		if !target.Living.Alive {
			continue
		}

		rs.draw(&target)
	}

	// Display images
    glfw.PollEvents()
    rs.Window.SwapBuffers()
}

func (rs RenderSystem) draw(node *RenderNode) {
	node.Display.X = node.Position.X
	node.Display.Y = node.Position.Y
	node.Display.Rotation = node.Position.Rotation

	// Create Vertex Array Object
	vao := MakeVao(node.Display.Points)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(node.Display.Points) / 3))
}

///////////////////
// LIVING SYSTEM //
///////////////////
type LivingSystem struct {
	Targets [][]LivingNode
}

func (ls LivingSystem) Update(time float32) {
	for x := range ls.Targets {
		for _, node := range ls.Targets[x]{
			ls.updateNodeState(&node)
		}
	}
}

// Updates the living state of the given node.
func (ls LivingSystem) updateNodeState(node *LivingNode) {
	node.Living.Alive = node.Living.AliveNext
	node.Living.AliveNext = node.Living.Alive

	liveCount := ls.liveNeighbors(node)

	// Rules of life
	if node.Living.Alive {
		// 1. Any live cell with fewer than two live neighbors dies, as if caused by underpopulation.
		if liveCount < 2 {
			node.Living.AliveNext = false
		}
	
		// 2. Any live cell with two or three live neighbors lives.
		if liveCount == 2 || liveCount == 3 {
			node.Living.AliveNext = true
		}
	
		// 3. Any live cell with more than three neighbors dies, as if by overpopulation.
		if liveCount > 3 {
			node.Living.AliveNext = false
		}
	} else {
		// 4. Any dead cell with exactly 3 neighbors becomes a live cell, as if by reproduction.
		if liveCount == 3 {
			node.Living.AliveNext = true
		}
	}
}

// Counts the number of living nodes around the given node.
func (ls LivingSystem) liveNeighbors(node *LivingNode) int {
	var liveCount int

	add := func(x, y int) {
		// Board edge checks
		if x == len(ls.Targets) {
			x = 0
		} else if x == -1 {
			x = len(ls.Targets) - 1
		}

		if y == len(ls.Targets[x]) {
			y = 0
		} else if y == -1 {
			y = len(ls.Targets[x]) - 1
		}

		if ls.Targets[x][y].Living.Alive {
			liveCount++
		}
	}

	x := int(node.Position.X)
	y := int(node.Position.Y)

	add(x - 1, y)   // To the left
    add(x + 1, y)   // To the right
    add(x, y + 1)   // up
    add(x, y - 1)   // down
    add(x - 1, y + 1) // top-left
    add(x + 1, y + 1) // top-right
    add(x - 1, y - 1) // bottom-left
    add(x + 1, y - 1) // bottom-right
    
    return liveCount
}

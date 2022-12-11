package engine

type Engine struct {
	Entities []*Entity
	RenderSystem *RenderSystem
	LivingSystem *LivingSystem
	RenderNodes []*RenderNode
	LivingNodes []*LivingNode
}

func NewEngine() (*Engine, error) {
	engine := &Engine{
		Entities: make([]*Entity, 0),

		RenderNodes: make([]*RenderNode, 0),
		LivingNodes: make([]*LivingNode, 0),
	}

	return engine, nil
}

func (e *Engine) AddRenderSystem(system *RenderSystem) {
	e.RenderSystem = system

	NodeAdded.AddListener(e.RenderSystem)
}

func (e *Engine) AddLivingSystem(system *LivingSystem) {
	e.LivingSystem = system

	NodeAdded.AddListener(e.LivingSystem)
}

// Add an entity to the simulation
func (e *Engine) AddEntity(entity *Entity) {
	_ = append(e.Entities, entity)

	renderNode := &RenderNode{ 
        Position: entity.Position,
        Display: entity.Display,
        Living: entity.Living,
    }

	livingNode := &LivingNode{
        Position: entity.Position,
        Living: entity.Living,
    }

	// Add nodes to lists
	e.RenderNodes = append(e.RenderNodes, renderNode)
	e.LivingNodes = append(e.LivingNodes, livingNode)

	// Notify Systems
	NodeAdded.Invoke(NodeAddedPayload{
		Class: Render,
		RenderNode: renderNode,
	})

	NodeAdded.Invoke(NodeAddedPayload{
		Class: Living,
		LivingNode: livingNode,
	})
}

// Remove an entity from the simulation
func (e *Engine) RemoveEntity(entity *Entity) {	
	for i, ent := range e.Entities {
		if ent == entity {
			_ = append(e.Entities[:i], e.Entities[i+1:]...)
		}
	}

	// Remove nodes?
}

func (e Engine) Update(time float32) {
	e.LivingSystem.Update(0)
	e.RenderSystem.Update(0)
}
package engine

import ( "container/list")

type Engine struct {
	Entities []*Entity
	RenderSystem *RenderSystem
	LivingSystem *LivingSystem
	Nodes map[string]*list.List
}

// Add an entity to the simulation
func (e Engine) AddEntity(entity *Entity) {
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

	e.Nodes["render"].PushBack(renderNode)
	e.Nodes["living"].PushBack(livingNode)
}

// Remove an entity from the simulation
func (e Engine) RemoveEntity(entity *Entity) {	
	for i, ent := range e.Entities {
		if ent == entity {
			_ = append(e.Entities[:i], e.Entities[i+1:]...)
		}
	}

	// Remove nodes?
}

func (e Engine) Update(time float32) {
	// TODO
}
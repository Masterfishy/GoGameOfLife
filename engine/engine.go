package engine

import ( "container/list")

type Engine struct {
	Entities []*Entity
	Systems []any
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
}

func (e Engine) AddSystem(system *any) {
	// TODO
}

func (e Engine) RemoveSystem(system *any) {
	// TODO
}

func (e Engine) GetNodeList(nodeClass string) list.List {
	// TODO
	var list list.List
	return list
}

func (e Engine) Update(time float32) {
	// TODO
}
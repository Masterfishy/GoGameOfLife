package engine

import ( "container/list")

type Engine struct {
	Entities list.List
	Systems []any
	Nodes map[string]list.List
}

func (e Engine) AddEntity(entity Entity) {
	// TODO
}

func (e Engine) RemoveEntity(entity Entity) {
	// TODO
}

func (e Engine) AddSystem(system any) {
	// TODO
}

func (e Engine) RemoveSystem(system any) {
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
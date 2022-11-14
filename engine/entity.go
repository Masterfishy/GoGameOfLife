package engine

type Entity struct {
	Name string
	Components map[string]any
}

func (e Entity) AddComponent(name string, component any) {
	e.Components[name] = component
}

func (e Entity) RemoveComponent(name string) {
	delete(e.Components, name)
}

func (e Entity) GetComponent(name string) any {
	return e.Components[name]
}

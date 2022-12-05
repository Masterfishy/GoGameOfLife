package engine

type Entity struct {
	Name string
	Position *PositionComponent
	Display *DisplayComponent
	Living *LivingComponent
}

func (e Entity) AddComponent(name string, component any) {
	// TODO if generic
}

func (e Entity) RemoveComponent(name string) {
	// TODO if generic
}

func (e Entity) GetComponent(name string) any {
	// TODO if generic
	return nil
}

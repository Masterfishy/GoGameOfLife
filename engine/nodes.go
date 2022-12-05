package engine

type RenderNode struct {
	Position *PositionComponent
	Display *DisplayComponent
	Living *LivingComponent
}

// Creates a new render node
func NewRenderNode(position *PositionComponent, display *DisplayComponent, living *LivingComponent) (*RenderNode, error) {
	node := &RenderNode{
		Position: position,
		Display: display,
		Living: living,
	}

	return node, nil
}

type LivingNode struct {
	Living *LivingComponent
	Position *PositionComponent
}

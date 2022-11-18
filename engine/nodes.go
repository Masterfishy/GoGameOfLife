package engine

type RenderNode struct {
	Position *PositionComponent
	Display *DisplayComponent
	Living *LivingComponent
}

type LivingNode struct {
	Living *LivingComponent
	Position *PositionComponent
}

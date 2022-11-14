package engine

type MoveNode struct {
	Position *PositionComponent
	Velocity *VelocityComponent
}

type RenderNode struct {
	Position *PositionComponent
	Display *DisplayComponent
	Living *LivingComponent
}

type LivingNode struct {
	Living *LivingComponent
	Position *PositionComponent
}

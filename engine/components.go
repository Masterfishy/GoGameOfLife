package engine

type PositionComponent struct {
	X float32
	Y float32
	Rotation float32
}

type DisplayComponent struct {
	Points []float32
	X float32
	Y float32
	Rotation float32
}

type LivingComponent struct {
	Alive bool
	AliveNext bool
}

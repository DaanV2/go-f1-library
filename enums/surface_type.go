package enums

// These types are from physics data and show what type of contact each wheel is experiencing.
type SurfaceType int8

const (
	Tarmac      SurfaceType = 0
	RumbleStrip SurfaceType = 1
	Concrete    SurfaceType = 2
	Rock        SurfaceType = 3
	Gravel      SurfaceType = 4
	Mud         SurfaceType = 5
	Sand        SurfaceType = 6
	Grass       SurfaceType = 7
	Water       SurfaceType = 8
	Cobblestone SurfaceType = 9
	Metal       SurfaceType = 10
	Ridged      SurfaceType = 11
)

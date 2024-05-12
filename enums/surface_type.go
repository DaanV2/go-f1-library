package enums

// These types are from physics data and show what type of contact each wheel is experiencing.
type SurfaceType int8

const (
	ST_Tarmac      SurfaceType = 0 // surface type: Tarmac
	ST_RumbleStrip SurfaceType = 1 // surface type: RumbleStrip
	ST_Concrete    SurfaceType = 2 // surface type: Concrete
	ST_Rock        SurfaceType = 3 // surface type: Rock
	ST_Gravel      SurfaceType = 4 // surface type: Gravel
	ST_Mud         SurfaceType = 5 // surface type: Mud
	ST_Sand        SurfaceType = 6 // surface type: Sand
	ST_Grass       SurfaceType = 7 // surface type: Grass
	ST_Water       SurfaceType = 8 // surface type: Water
	ST_Cobblestone SurfaceType = 9 // surface type: Cobblestone
	ST_Metal       SurfaceType = 10 // surface type: Metal
	ST_Ridged      SurfaceType = 11 // surface type: Ridged
)

package gamemap

// mapLayer is a 2D array containing layer information
type mapLayer [][]string

// GameMap is the structure for storing map information of the game world.
// mapWidth - map dimensions
// mapHeight - map dimensions
// mapLayerList - named list mapLayer's
type GameMap struct {
	mapWidth  int
	mapHeight int

	// GameMap layers
	mapLayerList map[string]mapLayer
}

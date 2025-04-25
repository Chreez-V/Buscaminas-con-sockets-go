package game

type Cell struct {
	HasMine   bool
	Revealed  bool
	Flagged   bool
	AdjacentMines int
}

package round2

type Cell struct {
	wasAlive             bool
	isAlive              bool
	aliveNeighboursCount int
	border               Border
}

func createCells(states [][]bool) [][]Cell {
	cells := make([][]Cell, len(states))
	for i := 0; i < len(states); i++ {
		cells[i] = make([]Cell, len(states[i]))
		for j := 0; j < len(states[i]); j++ {
			cells[i][j] = Cell{
				wasAlive: false,
				isAlive: states[i][j],
				aliveNeighboursCount: 0,
				border: nil,
			}
			if i == 0 || i == len(states) - 1 || j == 0 || j == len(states[i]) - 1 {
				cells[i][j].border = &EdgeBorder{}
			} else {
				cells[i][j].border = &FullBorder{}
			}
		}
	}
	return cells
}

type Border interface {
	visitNeighbours(cells [][]Cell, x int, y int, visitor func(neighbour *Cell))
}

type EdgeBorder struct {
}

func (border *EdgeBorder) visitNeighbours(cells [][]Cell, x int, y int, visitor func(neighbour *Cell)) {
	rightEdge := len(cells)
	bottomEdge := len(cells[0])
	for i := x - 1; i <= x + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if i == x && j == y {
				continue // skip myself
			}
			if i < 0 {
				// left edge
				continue
			}
			if j < 0 {
				// top edge
				continue
			}
			if i >= rightEdge {
				continue
			}
			if j >= bottomEdge {
				continue
			}
			visitor(&cells[i][j])
		}
	}
}

type FullBorder struct {
}

func (border *FullBorder) visitNeighbours(cells [][]Cell, x int, y int, visitor func(neighbour *Cell)) {
	visitor(&cells[x - 1][y - 1])
	visitor(&cells[x - 1][y])
	visitor(&cells[x - 1][y + 1])
	visitor(&cells[x][y - 1])
	visitor(&cells[x][y + 1])
	visitor(&cells[x + 1][y - 1])
	visitor(&cells[x + 1][y])
	visitor(&cells[x + 1][y + 1])
}

func (cell *Cell) justRevived() bool {
	return !cell.wasAlive && cell.isAlive
}

func (cell *Cell) justDied() bool {
	return cell.wasAlive && !cell.isAlive
}

func (cell *Cell) notifyThereIsNeighbourRevived() {
	cell.aliveNeighboursCount++
}

func (cell *Cell) notifyThereIsNeighbourDied() {
	cell.aliveNeighboursCount--
}

func (cell *Cell) updateIsAlive() {
	if cell.aliveNeighboursCount == 3 {
		cell.isAlive = true
	} else if cell.aliveNeighboursCount == 2 {
		cell.isAlive = cell.wasAlive
	} else {
		cell.isAlive = false
	}
}

func notifyNeighbours(cells [][]Cell, x int, y int) {
	myself := &cells[x][y]
	if myself.justRevived() {
		myself.wasAlive = true
		myself.border.visitNeighbours(cells, x, y, func(neighbour *Cell) {
			neighbour.notifyThereIsNeighbourRevived()
		})
	} else if myself.justDied() {
		myself.wasAlive = false
		myself.border.visitNeighbours(cells, x, y, func(neighbour *Cell) {
			neighbour.notifyThereIsNeighbourDied()
		})
	}
}

func runOneCycle(cells [][]Cell) {
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			notifyNeighbours(cells, i, j)
		}
	}
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			cells[i][j].updateIsAlive()
		}
	}
}
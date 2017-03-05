package round1

func runOneCycle(input [][]bool, output [][]bool, x int, y int) {
	aliveNeighbours := countAliveNeighbours(input, x, y)
	if aliveNeighbours == 3 {
		output[x][y] = true
	} else if aliveNeighbours == 2 {
		output[x][y] = input[x][y]
	} else {
		output[x][y] = false
	}
}

func countAliveNeighbours(input [][]bool, x int, y int) int {
	count := 0
	rightEdge := len(input)
	bottomEdge := len(input[0])
	for i := x -1; i <= x +1; i++ {
		for j := y-1; j <= y+1; j++ {
			if i == x && j == y {
				continue // skip myself
			}
			if i < 0 { // left edge
				continue
			}
			if j < 0 { // top edge
				continue
			}
			if i >= rightEdge {
				continue
			}
			if j >= bottomEdge {
				continue
			}
			if input[i][j] {
				count++
			}
		}
	}
	return count
}
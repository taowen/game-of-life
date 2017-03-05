package round2

import (
	"testing"
	"github.com/taowen/game-of-life/require"
	"fmt"
)

func Test_notifyNeighbours(t *testing.T) {
	t.Run("central revived", func(t *testing.T) {
		should := require.New(t)
		cells := createCells([][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		})
		notifyNeighbours(cells, 1, 1)
		should.Equal(1, cells[0][0].aliveNeighboursCount)
		should.Equal(1, cells[2][2].aliveNeighboursCount)
		should.Equal(0, cells[1][1].aliveNeighboursCount)
	})
	t.Run("central revived=>died", func(t *testing.T) {
		should := require.New(t)
		cells := createCells([][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		})
		notifyNeighbours(cells, 1, 1)
		should.Equal(1, cells[0][0].aliveNeighboursCount)
		should.Equal(1, cells[2][2].aliveNeighboursCount)
		cells[1][1].isAlive = false
		notifyNeighbours(cells, 1, 1)
		should.Equal(0, cells[0][0].aliveNeighboursCount)
		should.Equal(0, cells[2][2].aliveNeighboursCount)
	})
	t.Run("edge", func(t *testing.T) {
		should := require.New(t)
		cells := createCells([][]bool{
			{false, true, false},
			{false, false, false},
			{false, false, false},
		})
		notifyNeighbours(cells, 0, 1)
		should.Equal(1, cells[0][0].aliveNeighboursCount)
		should.Equal(1, cells[1][1].aliveNeighboursCount)
		should.Equal(0, cells[2][2].aliveNeighboursCount)
	})
}

func Test_updateIsAlive(t *testing.T) {
	t.Run("== 3", func(t *testing.T) {
		should := require.New(t)
		cell := Cell{}
		cell.aliveNeighboursCount = 3
		cell.updateIsAlive()
		should.True(cell.isAlive)
	})
	t.Run("alive == 2 alive", func(t *testing.T) {
		should := require.New(t)
		cell := Cell{}
		cell.wasAlive = true
		cell.aliveNeighboursCount = 2
		cell.updateIsAlive()
		should.True(cell.isAlive)
	})
	t.Run("dead == 2 dead", func(t *testing.T) {
		should := require.New(t)
		cell := Cell{}
		cell.wasAlive = false
		cell.aliveNeighboursCount = 2
		cell.updateIsAlive()
		should.False(cell.isAlive)
	})
	t.Run("otherwise, dead", func(t *testing.T) {
		should := require.New(t)
		cell := Cell{}
		cell.wasAlive = true
		cell.isAlive = true
		cell.aliveNeighboursCount = 1
		cell.updateIsAlive()
		should.False(cell.isAlive)
	})
}

func Test_runOneCycle(t *testing.T) {
	should := require.New(t)
	cells := createCells([][]bool{
		{false, false, false},
		{false, false, false},
		{true, true, true},
	})
	runOneCycle(cells)
	expect := [][]bool {
		{false, false, false},
		{false, true, false},
		{false, true, false},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			should.Equal(expect[i][j], cells[i][j].isAlive, fmt.Sprintf("%v,%v", i, j))
		}
	}
}
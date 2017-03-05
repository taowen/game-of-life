package round1

import (
	"testing"
	"github.com/taowen/game-of-life/require"
)

func Test_runOneCycle(t *testing.T) {
	output := [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
	t.Run("==3", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, false, false},
			{false, false, false},
			{true, true, true},
		}
		runOneCycle(input, output, 1, 1)
		should.True(output[1][1])
	})
	t.Run("dead ==2 dead", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, true, true},
		}
		runOneCycle(input, output, 1, 1)
		should.False(output[1][1])
	})
	t.Run("alive ==2 alive", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, false, false},
			{false, true, false},
			{false, true, true},
		}
		runOneCycle(input, output, 1, 1)
		should.True(output[1][1])
	})
	t.Run("otherwise, dead", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, false, false},
			{false, true, false},
			{false, true, false},
		}
		runOneCycle(input, output, 1, 1)
		should.False(output[1][1])
	})
	t.Run("top edge", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, true, false},
			{false, false, false},
			{false, false, false},
		}
		runOneCycle(input, output, 0, 1)
		should.False(output[0][1])
	})
	t.Run("left edge", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{true, false, false},
			{false, false, false},
			{false, false, false},
		}
		output := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
		runOneCycle(input, output, 0, 0)
		should.False(output[0][0])
	})
}

func Test_countAliveNeighbours(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
		}
		expect := [][]int{
			{3, 5, 3},
			{5, 8, 5},
			{3, 5, 3},
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				should.Equal(expect[i][j], countAliveNeighbours(input, i, j))
			}
		}
	})
	t.Run("all false", func(t *testing.T) {
		should := require.New(t)
		input := [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
		expect := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				should.Equal(expect[i][j], countAliveNeighbours(input, i, j))
			}
		}
	})
}


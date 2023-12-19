package maze

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMazeInit(t *testing.T) {
	maze := NewMaze(3, 4)
	expectedType := reflect.TypeOf(&Maze{})
	actualType := reflect.TypeOf(maze)
	assert.Equal(t, expectedType, actualType)
	assert.Equal(t, 4, maze.Width)
	assert.Equal(t, 3, maze.Height)
}

func TestMazeSetter(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		mazeX  int
		mazeY  int
		target string
	}{
		{x: 0, y: 1, mazeX: 2, mazeY: 2, target: "Start"},
		{x: 2, y: 1, mazeX: 3, mazeY: 2, target: "Goal"},
	}
	var gotX int
	var gotY int
	var err error
	for _, tc := range tests {
		maze := NewMaze(tc.mazeY, tc.mazeX)
		switch tc.target {
		case "Start":
			maze.SetStart(tc.x, tc.y)
			gotX, gotY, err = maze.GetStart()
			assert.Nil(t, err)
			assert.Equal(t, tc.x, gotX)
			assert.Equal(t, tc.y, gotY)
		case "Goal":
			maze.SetGoal(tc.x, tc.y)
			gotX, gotY, err = maze.GetGoal()
			assert.Nil(t, err)
			assert.Equal(t, tc.x, gotX)
			assert.Equal(t, tc.y, gotY)
		}
	}
}

func TestMazeExistPath(t *testing.T) {
	tests := []struct {
		mazeX     int
		mazeY     int
		startX    int
		startY    int
		endX      int
		endY      int
		obstacles [][]int
		reached   bool
	}{
		{
			mazeX:  3,
			mazeY:  3,
			startX: 0,
			startY: 0,
			endX:   2,
			endY:   2,
			obstacles: [][]int{
				{1, 1},
				{0, 1},
			},
			reached: true,
		},
		{
			mazeX:  3,
			mazeY:  3,
			startX: 2,
			startY: 2,
			endX:   0,
			endY:   0,
			obstacles: [][]int{
				{1, 1},
				{0, 1},
			},
			reached: true,
		},
		{
			mazeX:  3,
			mazeY:  3,
			startX: 0,
			startY: 0,
			endX:   2,
			endY:   2,
			obstacles: [][]int{
				{0, 1},
				{1, 1},
				{2, 1},
			},
			reached: false,
		},
	}
	for _, tc := range tests {
		maze := NewMaze(tc.mazeY, tc.mazeX)
		for _, obstacle := range tc.obstacles {
			maze.SetObstacle(obstacle[0], obstacle[1])
		}
		maze.SetStart(tc.startX, tc.startY)
		maze.SetGoal(tc.endX, tc.endY)
		reached := maze.ExistPath()
		assert.Equal(t, tc.reached, reached)
	}
}

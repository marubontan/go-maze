package main

import (
	"fmt"

	"github.com/marubontan/go-maze"
)

func main() {
	maze := maze.NewMaze(3, 3)
	maze.SetStart(0, 0)
	maze.SetGoal(2, 2)
	maze.SetObstacle(1, 1)
	maze.SetObstacle(0, 1)
	maze.Print()
	fmt.Printf("Path from start to goal: %t\n", maze.ExistPath())
}

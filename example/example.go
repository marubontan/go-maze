package main

import (
	"fmt"

	maze "github.com/marubontan/go-maze/maze"
)

func main() {
	maze := maze.NewMaze(3, 3)
	maze.SetStart(0, 0)
	maze.SetGoal(2, 2)
	maze.SetObstacle(1, 1)
	maze.SetObstacle(0, 1)
	maze.SetTrap(2, 1)
	maze.SetGoalReward(1.0)
	maze.SetTrapPenalty(-0.1)
	maze.Print()
	fmt.Printf("Path from start to goal: %t\n", maze.ExistPath())
	fmt.Printf("Goal Reward: %f\n", *maze.GoalReward)
	fmt.Printf("Trap Penalty: %f\n", *maze.TrapPenalty)
}

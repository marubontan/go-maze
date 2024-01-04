# go-maze [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## Overview
`go-maze` is simple maze making library, expecting to be used for some algorithm and reinforcememnt leaning experiments.

## How to install
```
go get github.com/marubontan/go-maze
```

## How to use
### Actions
Right: 0
Down: 1
Left: 2
Up: 3

### Sample Code

```Go
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
	maze.Reset()
	nextState, isGoal, reward, _ := maze.Step(0)
	fmt.Printf("action: 0, NextState: %v, isGoal: %t, reward: %f\n", nextState, isGoal, reward)
	nextState, isGoal, reward, _ = maze.Step(0)
	fmt.Printf("action: 0, NextState: %v, isGoal: %t, reward: %f\n", nextState, isGoal, reward)
	nextState, isGoal, reward, _ = maze.Step(1)
	fmt.Printf("action: 1, NextState: %v, isGoal: %t, reward: %f\n", nextState, isGoal, reward)
	nextState, isGoal, reward, _ = maze.Step(1)
	fmt.Printf("action: 1, NextState: %v, isGoal: %t, reward: %f\n", nextState, isGoal, reward)
}

```
The code above prints.  
```
 - - -
|S| | |
 - - -
|X|X|T|
 - - -
| | |G|
 - - -
Path from start to goal: true
Goal Reward: 1.000000
Trap Penalty: -0.100000
action: 0, NextState: [1 0], isGoal: false, reward: 0.000000
action: 0, NextState: [2 0], isGoal: false, reward: 0.000000
action: 1, NextState: [2 1], isGoal: false, reward: -0.100000
action: 1, NextState: [2 2], isGoal: true, reward: 1.000000
```
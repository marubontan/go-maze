# go-maze [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## Overview
`go-maze` is simple maze making library, expecting to be used for some algorithm and reinformemnt leaning experiments.

## How to install
```
go get github.com/marubontan/go-maze
```

## How to use
```Go
package main

import (
	"fmt"

	maze "github.com/marubontan/go-maze/maze"
)

func main() {
	playground := maze.NewMaze(3, 3)
	playground.SetStart(0, 0)
	playground.SetGoal(2, 2)
	playground.SetObstacle(1, 1)
	playground.Print()
	fmt.Printf("The path from start to end %v", playground.ExistPath())
}
```
The code above prints.  
```
 - - -
|S| | |
 - - -
| |X| |
 - - -
| | |G|
 - - -
The path from start to end true
```
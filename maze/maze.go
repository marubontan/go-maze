package maze

import (
	"errors"
	"fmt"
)

const (
	Empty = iota
	Obstacle
	Trap
	Start
	Goal
)

const (
	Right = iota
	Down
	Left
	Up
)

type Block struct {
	BlockType  int
	Attributes map[string]any
}
type Maze struct {
	Height      int
	Width       int
	Blocks      [][]Block
	GoalReward  *float64
	TrapPenalty *float64
	State       *[2]int
}

func NewMaze(h, w int) *Maze {
	m := &Maze{
		Height: h,
		Width:  w,
		Blocks: make([][]Block, h),
	}
	for i := 0; i < h; i++ {
		m.Blocks[i] = make([]Block, w)
	}
	return m
}

func (m *Maze) SetStart(x, y int) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("invalid start position")
	}
	m.Blocks[y][x] = Block{BlockType: Start}
	return nil
}

func (m *Maze) GetStart() (int, int, error) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.Blocks[y][x].BlockType == Start {
				return x, y, nil
			}
		}
	}
	return -1, -1, errors.New("start not found")
}

func (m *Maze) SetGoal(x, y int) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("invalid goal position")
	}
	m.Blocks[y][x] = Block{BlockType: Goal}
	return nil
}

func (m *Maze) SetGoalReward(reward float64) {
	m.GoalReward = &reward
}

func (m *Maze) SetTrapPenalty(penalty float64) {
	m.TrapPenalty = &penalty
}

func (m *Maze) GetGoal() (int, int, error) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.Blocks[y][x].BlockType == Goal {
				return x, y, nil
			}
		}
	}
	return -1, -1, errors.New("goal not found")
}

func (m *Maze) SetObstacle(x, y int) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("invalid goal position")
	}
	m.Blocks[y][x] = Block{BlockType: Obstacle}
	return nil
}

func (m *Maze) SetTrap(x, y int) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("invalid goal position")
	}
	m.Blocks[y][x] = Block{BlockType: Trap}
	return nil
}

func (m *Maze) Step(action int) ([2]int, bool, float64, error) {
	var nextStateCandidate [2]int
	switch action {
	case Right:
		nextStateCandidate = [2]int{m.State[0] + 1, m.State[1]}
	case Down:
		nextStateCandidate = [2]int{m.State[0], m.State[1] + 1}
	case Left:
		nextStateCandidate = [2]int{m.State[0] - 1, m.State[1]}
	case Up:
		nextStateCandidate = [2]int{m.State[0], m.State[1] - 1}
	}
	if !m.IsAvailable(nextStateCandidate[0], nextStateCandidate[1]) {
		return *m.State, false, 0, nil
	}
	reward := 0.0
	isGoal := false
	if m.Blocks[nextStateCandidate[1]][nextStateCandidate[0]].BlockType == Goal {
		reward = *m.GoalReward
		isGoal = true
	} else if m.Blocks[nextStateCandidate[1]][nextStateCandidate[0]].BlockType == Trap {
		reward = *m.TrapPenalty
	}
	m.State = &nextStateCandidate
	return nextStateCandidate, isGoal, reward, nil
}

func (m *Maze) Reset() {
	x, y, err := m.GetStart()
	if err != nil {
		x, y = 0, 0
		m.SetStart(x, y)
	}
	m.State = &[2]int{x, y}
}

func (m *Maze) IsAvailable(x, y int) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}
	if m.Blocks[y][x].BlockType == Obstacle {
		return false
	}
	return true
}

func drawHorizontalLine(w int) {
	for i := 0; i < w; i++ {
		fmt.Print(" -")
	}
	fmt.Print("\n")
}

func (m *Maze) Print() {
	drawHorizontalLine(m.Width)
	for y := 0; y < m.Height; y++ {
		fmt.Print("|")
		for x := 0; x < m.Width; x++ {
			switch m.Blocks[y][x].BlockType {
			case Empty:
				fmt.Print(" |")
			case Obstacle:
				fmt.Print("X|")
			case Trap:
				fmt.Print("T|")
			case Start:
				fmt.Print("S|")
			case Goal:
				fmt.Print("G|")
			}
		}
		fmt.Print("\n")
		drawHorizontalLine(m.Width)
	}
}

func (m *Maze) isAvailable(x, y int, seen [][]bool) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return false
	}
	if m.Blocks[y][x].BlockType == Obstacle {
		return false
	}
	if seen[y][x] {
		return false
	}
	return true
}

func (m *Maze) dfs(x, y, endX, endY int, seen [][]bool) bool {
	if !m.isAvailable(x, y, seen) {
		return false
	}
	seen[y][x] = true
	if x == endX && y == endY {
		return true
	}
	if m.dfs(x+1, y, endX, endY, seen) || m.dfs(x, y+1, endX, endY, seen) || m.dfs(x-1, y, endX, endY, seen) || m.dfs(x, y-1, endX, endY, seen) {
		return true
	}
	return false
}

func (m *Maze) ExistPath() bool {
	startX, startY, error := m.GetStart()
	if error != nil {
		return false
	}
	endX, endY, error := m.GetGoal()
	if error != nil {
		return false
	}

	seen := make([][]bool, m.Height)
	for i := 0; i < m.Height; i++ {
		seen[i] = make([]bool, m.Width)
	}

	return m.dfs(startX, startY, endX, endY, seen)
}

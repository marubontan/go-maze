package maze

import "fmt"

const (
	Empty = iota
	Obstacle
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
	blockType int
}
type Maze struct {
	height int
	width  int
	blocks [][]Block
}

func NewMaze(h, w int) *Maze {
	m := &Maze{
		height: h,
		width:  w,
		blocks: make([][]Block, h),
	}
	for i := 0; i < h; i++ {
		m.blocks[i] = make([]Block, w)
	}
	return m
}

func (m *Maze) SetStart(x, y int) {
	m.blocks[y][x] = Block{blockType: Start}
}

func (m *Maze) GetStart() (int, int) {
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if m.blocks[y][x].blockType == Start {
				return x, y
			}
		}
	}
	return -1, -1
}

func (m *Maze) SetGoal(x, y int) {
	m.blocks[y][x] = Block{blockType: Goal}
}

func (m *Maze) SetObstacle(x, y int) {
	m.blocks[y][x] = Block{blockType: Obstacle}
}

func drawHorizontalLine(w int) {
	for i := 0; i < w; i++ {
		fmt.Print(" -")
	}
	fmt.Print("\n")
}

func (m *Maze) Print() {
	drawHorizontalLine(m.width)
	for y := 0; y < m.height; y++ {
		fmt.Print("|")
		for x := 0; x < m.width; x++ {
			switch m.blocks[y][x].blockType {
			case Empty:
				fmt.Print(" |")
			case Obstacle:
				fmt.Print("X|")
			case Start:
				fmt.Print("S|")
			case Goal:
				fmt.Print("G|")
			}
		}
		fmt.Print("\n")
		drawHorizontalLine(m.width)
	}
}

func (m *Maze) isAvailable(x, y int, seen [][]bool) bool {
	if x < 0 || x >= m.width || y < 0 || y >= m.height {
		return false
	}
	if m.blocks[y][x].blockType == Obstacle {
		return false
	}
	if seen[y][x] {
		return false
	}
	return true
}

func (m *Maze) isGoal(x, y int) bool {
	return m.blocks[y][x].blockType == Goal
}

func (m *Maze) dfs(x, y int, seen [][]bool) bool {
	if !m.isAvailable(x, y, seen) {
		return false
	}
	seen[y][x] = true
	if m.isGoal(x, y) {
		return true
	}
	if m.dfs(x+1, y, seen) || m.dfs(x, y+1, seen) || m.dfs(x-1, y, seen) || m.dfs(x, y-1, seen) {
		return true
	}
	return false
}

func (m *Maze) ExistPath() bool {
	startX, startY := m.GetStart()
	if startX == -1 || startY == -1 {
		return false
	}
	seen := make([][]bool, m.height)
	for i := 0; i < m.height; i++ {
		seen[i] = make([]bool, m.width)
	}

	return m.dfs(startX, startY, seen)
}

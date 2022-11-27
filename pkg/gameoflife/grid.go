package gameoflife

import "fmt"

type Cell struct {
	x int64
	y int64
}

type coordinate string

func (c Cell) getCoordinate() coordinate {
	return getCoordinate(c.x, c.y)
}

func getCoordinate(x, y int64) coordinate {
	return coordinate(fmt.Sprintf("%d:%d", x, y))
}

type Grid struct {
	aliveCells map[coordinate]Cell
}

func NewGrid(seed ...Cell) Grid {
	grid := Grid{aliveCells: make(map[coordinate]Cell, len(seed))}
	for _, cell := range seed {
		grid.aliveCells[cell.getCoordinate()] = cell
	}
	return grid
}

func (g Grid) GetAliveCells() []Cell {
	aliveCells := make([]Cell, 0, len(g.aliveCells))
	for _, cell := range g.aliveCells {
		aliveCells = append(aliveCells, cell)
	}
	return aliveCells
}

func (g Grid) Tick() {
	cellsToGiveBirth := make([]Cell, 0)
	cellsToKill := make([]Cell, 0)

	// check which cells to kill and which to give birth
	for _, cell := range g.aliveCells {
		aliveNeighbours, deadNeighbours := g.getNeighbours(cell)

		// overpopulation or loneliness
		if len(aliveNeighbours) > 3 || len(aliveNeighbours) < 2 {
			cellsToKill = append(cellsToKill, cell)
		}

		// new cells
		for _, deadNeighbourCell := range deadNeighbours {
			deadNeighbourCellAliveNeighbours, _ := g.getNeighbours(deadNeighbourCell)
			if len(deadNeighbourCellAliveNeighbours) == 3 {
				cellsToGiveBirth = append(cellsToGiveBirth, deadNeighbourCell)
			}
		}
	}

	// kill and give birth
	for _, cellToKill := range cellsToKill {
		g.kill(cellToKill.x, cellToKill.y)
	}

	for _, cellToGiveBirth := range cellsToGiveBirth {
		g.giveBirth(cellToGiveBirth.x, cellToGiveBirth.y)
	}
}

func (g Grid) getNeighbours(cell Cell) (alive []Cell, dead []Cell) {
	alive = make([]Cell, 0)
	dead = make([]Cell, 0)

	// FIXME check max/min int64
	for x := cell.x - 1; x <= cell.x+1; x++ {
		for y := cell.y - 1; y <= cell.y+1; y++ {
			if g.isSame(x, y, cell) {
				continue
			}
			if g.isAlive(x, y) {
				alive = append(alive, Cell{x, y})
			} else {
				dead = append(dead, Cell{x, y})
			}
		}
	}

	return
}

func (g Grid) isAlive(x, y int64) bool {
	_, ok := g.aliveCells[getCoordinate(x, y)]
	return ok
}

func (g Grid) isSame(x int64, y int64, cell Cell) bool {
	return cell.x == x && cell.y == y
}

func (g Grid) kill(x int64, y int64) {
	delete(g.aliveCells, getCoordinate(x, y))
}

func (g Grid) giveBirth(x int64, y int64) {
	g.aliveCells[getCoordinate(x, y)] = Cell{x, y}
}

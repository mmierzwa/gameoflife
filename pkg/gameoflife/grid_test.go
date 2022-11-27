package gameoflife

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type generations []Grid

func TestGrid(t *testing.T) {
	testCases := map[string]generations{
		"empty grid": []Grid{
			NewGrid(),
			NewGrid(),
		},
		"single cell": []Grid{
			NewGrid(Cell{0, 0}),
			NewGrid(),
		},
		"tree neighbours to survive and one to be born": []Grid{
			NewGrid(Cell{1, 0}, Cell{0, 0}, Cell{0, 1}),
			NewGrid(Cell{1, 0}, Cell{0, 0}, Cell{0, 1}, Cell{1, 1}),
		},
	}

	for testName, generations := range testCases {
		t.Run(testName, func(t *testing.T) {
			if len(generations) < 2 {
				t.Skip("not enough generations in test case")
			}

			for tickNumber := 1; tickNumber < len(generations); tickNumber++ {
				current := generations[tickNumber-1]

				current.Tick()

				assert.ElementsMatch(t, current.GetAliveCells(), generations[tickNumber].GetAliveCells())
			}
		})
	}
}

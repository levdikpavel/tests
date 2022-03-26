package equations

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSquareEquation(t *testing.T) {
	suite.Run(t, new(SquareEquationSuite))
}

type SquareEquationSuite struct {
	suite.Suite
}

func (s *SquareEquationSuite) TestNoSolution() {
	roots, err := solveSquareEquation(1,0,1)
	s.Require().NoError(err)
	s.Require().Empty(roots)
}

func (s *SquareEquationSuite) TestRoots() {
	roots, err := solveSquareEquation(1,0,-1)
	s.Require().NoError(err)
	s.Require().Len(roots, 2)
	s.assertRoots([]float64{-1, 1}, roots)
}

func (s *SquareEquationSuite) TestDoubleRoot() {
	roots, err := solveSquareEquation(1,2,1)
	s.Require().NoError(err)
	s.Require().Len(roots, 2)
	s.assertRoots([]float64{-1, -1}, roots)
}

func (s *SquareEquationSuite) TestNotSquare() {
	roots, err := solveSquareEquation(0,1,1)
	s.Require().Error(err)
	s.Require().Empty(roots)

	roots, err = solveSquareEquation(epsilon,1,1)
	s.Require().Error(err)
	s.Require().Empty(roots)
}

func (s *SquareEquationSuite) TestDoubleRootEpsilon() {
	diff := epsilon / 4
	roots, err := solveSquareEquation(1+diff, 2+2*diff, 1+diff)
	s.Require().NoError(err)
	s.Require().Len(roots, 2)
	s.assertRoots([]float64{-1, -1}, roots)
}

func (s *SquareEquationSuite) TestMaxInt() {
	max := float64(maxInt)
	roots, err := solveSquareEquation(max,2*max,max)
	s.Require().NoError(err)
	s.Require().Len(roots, 2)
	s.assertRoots([]float64{-1, -1}, roots)
}

func (s *SquareEquationSuite) TestNotNumber() {
	roots, err := solveSquareEquation(math.Inf(1), 1, 0)
	s.Require().Error(err)
	s.Require().Empty(roots)

	roots, err = solveSquareEquation(1, math.Inf(1), 0)
	s.Require().Error(err)
	s.Require().Empty(roots)

	roots, err = solveSquareEquation(1, 0, math.NaN())
	s.Require().Error(err)
	s.Require().Empty(roots)
}

func (s *SquareEquationSuite) assertRoots(expected, actual []float64) {
	sort.Slice(expected, func (i, j int) bool {
		return expected[i] < expected[j]
	})
	sort.Slice(actual, func (i, j int) bool {
		return actual[i] < actual[j]
	})
	s.Require().InEpsilonSlicef(expected, actual, epsilon, "expected: %v, actual: %v", expected, actual)
}

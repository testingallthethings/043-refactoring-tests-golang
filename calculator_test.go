package bowling_test

import (
	"bowling"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BowlingGameCalculatorSuite struct {
	suite.Suite
}

func TestBowlingGameCalculatorSuite(t *testing.T) {
	suite.Run(t, new(BowlingGameCalculatorSuite))
}

func (s *BowlingGameCalculatorSuite) TestKeepingScoreOfFirstBallInEveryFrame() {
	s.Equal(45, bowling.Score("-- 1- 2- 3- 4- 5- 6- 7- 8- 9-"))
}

func (s *BowlingGameCalculatorSuite) TestKeepingScoreOfSecondBallInEveryFrame() {
	s.Equal(45, bowling.Score("-- -1 -2 -3 -4 -5 -6 -7 -8 -9"))
}

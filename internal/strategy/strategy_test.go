package strategy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
)

func TestFixedStrategy(t *testing.T) {
	tests := []struct {
		score          int
		holdingValue   int
		expectedResult bool
	}{
		{
			score:          5,
			holdingValue:   5,
			expectedResult: true,
		},
		{
			score:          5,
			holdingValue:   10,
			expectedResult: false,
		},
	}

	for _, test := range tests {
		fixedStrategy := strategy.NewFixedStrategy(test.holdingValue)
		output := fixedStrategy.ShouldHold(test.score)
		want := test.expectedResult

		a := assert.New(t)
		a.Equal(output, want)
	}

}

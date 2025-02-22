package game_test

import (
	"testing"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/game"
	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	playingStratey := strategy.NewFixedStrategy(10)
	player := game.NewPlayer("abc", playingStratey)
	score := player.Roll()

	// set score
	player.AddScore(score)

	a := assert.New(t)
	a.Equal(score, player.TotalScore())

	// reset score
	player.ResetScore()
	a.Equal(0, player.TotalScore())

}

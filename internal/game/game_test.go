package game_test

import (
	"testing"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/game"
	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	a := assert.New(t)

	playingStratey := strategy.NewFixedStrategy(10)
	player1 := game.NewPlayer("abc", playingStratey)
	player2 := game.NewPlayer("sar", playingStratey)
	players := []*game.Player{player1, player2}
	game := game.NewGame(players, 100)
	a.NotNil(game.Play())
}

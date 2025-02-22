package simulator_test

import (
	"testing"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/game"
	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
	"github.com/devsarvesh92/game-of-pig-exercise/pkg/simulator"
)

func TestSimulator(t *testing.T) {
	fixedStrategy := strategy.NewFixedStrategy(10)
	player1 := game.NewPlayer("sarvesh", fixedStrategy)
	player2 := game.NewPlayer("pradnya", fixedStrategy)

	game := game.NewGame([]*game.Player{player1, player2}, 100)

	simulator := simulator.NewSimulator(10, game)
	simulator.RunGames()
}

package simulator

import (
	"fmt"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/game"
)

type Result struct {
	Winner string
	Score  int
}

type Simulator struct {
	numberOfRounds int
	game           *game.Game
}

func NewSimulator(numberOfGames int, game *game.Game) *Simulator {
	return &Simulator{
		numberOfRounds: numberOfGames,
		game:           game,
	}
}

func (s *Simulator) RunGames() {
	for i := 0; i < s.numberOfRounds; i++ {
		winner := s.game.Play()
		fmt.Println("Winner of game is", winner.Name, winner.Score)
	}
}

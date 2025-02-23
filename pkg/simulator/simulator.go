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
	player1Winnings := 0
	playersHolding := make([]int, 0, len(s.game.GetPlayers()))

	for _, player := range s.game.GetPlayers() {
		playersHolding = append(playersHolding, player.GetHoldingValue())
	}

	for i := 0; i < s.numberOfRounds; i++ {
		result := s.game.Play()
		if result {
			player1Winnings += 1
		}
	}
	gamesWon := player1Winnings
	winningPercentage := (float64(gamesWon) / float64(s.numberOfRounds)) * 100
	gamesLost := s.numberOfRounds - gamesWon
	losingPercentage := 100 - winningPercentage
	player1HolderingValue := playersHolding[0]
	player2HolderingValue := playersHolding[1]

	fmt.Printf("Holding at %v vs Holding at %v: wins %v (%v %%), looses: %v (%v %%)", player1HolderingValue, player2HolderingValue, gamesWon, winningPercentage, gamesLost, losingPercentage)
}

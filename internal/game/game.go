package game

import "fmt"

type Game struct {
	players []*Player
	target  int
}

func NewGame(players []*Player, target int) *Game {
	return &Game{
		players: players,
		target:  target,
	}
}

func (game Game) Play() *Player {

	// Reset score for all players
	for _, player := range game.players {
		player.ResetScore()
	}

	for {
		for _, player := range game.players {
			score := player.Roll()
			player.AddScore(score)
			if player.TotalScore() >= game.target {
				fmt.Println("Winner of game", player.name)
				return player
			}
		}
	}
}

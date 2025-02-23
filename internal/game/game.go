package game

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

func (game Game) Play() bool {

	// Reset score for all players
	for _, player := range game.players {
		player.ResetScore()
	}

	for {
		for _, player := range game.players {
			score := player.Roll()
			player.AddScore(score)
			if player.TotalScore() >= game.target {
				if player.Name == game.players[0].Name {
					return true
				} else {
					return false
				}
			}
		}
	}
}

func (game Game) GetPlayers() []*Player {
	return game.players
}

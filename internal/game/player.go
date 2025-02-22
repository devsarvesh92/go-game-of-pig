package game

import (
	"math/rand"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
)

type Player struct {
	score           int
	playingStrategy strategy.PlayingStrategy
}

func NewPlayer(playingStrategy strategy.PlayingStrategy) *Player {
	return &Player{playingStrategy: playingStrategy, score: 0}
}

func (player *Player) Roll() int {
	totalScore := 0
	for {
		roll := rand.Intn(6) + 1
		if roll == 1 {
			return totalScore
		}
		if player.playingStrategy.ShouldHold(player.score) {
			return totalScore
		}
		totalScore += roll
	}
}

func (player *Player) TotalScore() int {
	return player.score
}

func (player *Player) ResetScore() {
	player.score = 0
}

func (player *Player) AddScore(score int) {
	player.score += score
}

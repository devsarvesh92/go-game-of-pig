package game

import (
	"math/rand"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
)

type Player struct {
	Name            string
	Score           int
	playingStrategy strategy.PlayingStrategy
}

func NewPlayer(name string, playingStrategy strategy.PlayingStrategy) *Player {
	return &Player{Name: name, playingStrategy: playingStrategy, Score: 0}
}

func (player *Player) Roll() int {
	totalScore := 0
	for {
		roll := rand.Intn(6) + 1
		if roll == 1 {
			return totalScore
		}
		if player.playingStrategy.ShouldHold(totalScore) {
			return totalScore
		}
		totalScore += roll
	}
}

func (player *Player) TotalScore() int {
	return player.Score
}

func (player *Player) ResetScore() {
	player.Score = 0
}

func (player *Player) AddScore(score int) {
	player.Score += score
}

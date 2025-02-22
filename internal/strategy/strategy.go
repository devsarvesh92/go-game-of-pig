package strategy

type PlayingStrategy interface {
	ShouldHold(score int) bool
}

type FixedStrategy struct {
	holdingValue int
}

// Constructor
func NewFixedStrategy(holdingScore int) *FixedStrategy {
	return &FixedStrategy{holdingValue: holdingScore}
}

// Attach method to strcut
func (fixedStrategy FixedStrategy) ShouldHold(score int) bool {
	return score >= fixedStrategy.holdingValue
}

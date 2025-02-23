package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/devsarvesh92/game-of-pig-exercise/internal/game"
	"github.com/devsarvesh92/game-of-pig-exercise/internal/strategy"
	"github.com/devsarvesh92/game-of-pig-exercise/pkg/simulator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pig [playername1][num1] [playername2][num2]",
	Short: "A simple CLI that takes two numbers",
	Long: `This CLI application takes two numbers as positional arguments
and performs operations on them. Example: pig 10 15`,
	Args: cobra.ExactArgs(4), // Requires exactly 2 arguments
	Run: func(cmd *cobra.Command, args []string) {
		num1, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: first argument must be a number, got %s\n", args[0])
			os.Exit(1)

		}

		num2, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Error: first argument must be a number, got %s\n", args[1])
			os.Exit(1)

		}

		player1Name := args[0]
		player2Name := args[2]

		player1Strategy, player2Strategy := BuildStrategies(num1, num2)
		player1, player2 := BuildPlayers(player1Name, player1Strategy, player2Name, player2Strategy)
		game := BuildGame([]*game.Player{player1, player2})
		simulator := BuildSimulator(game)
		simulator.RunGames()
	},
}

func BuildStrategies(holdScore1 int, holderScore2 int) (*strategy.FixedStrategy, *strategy.FixedStrategy) {
	return strategy.NewFixedStrategy(holdScore1), strategy.NewFixedStrategy(holderScore2)
}

func BuildPlayers(player1Name string, player1Strategy *strategy.FixedStrategy, player2Name string, player2Strategy *strategy.FixedStrategy) (*game.Player, *game.Player) {
	return game.NewPlayer(player1Name, player1Strategy), game.NewPlayer(player2Name, player2Strategy)
}

func BuildGame(players []*game.Player) *game.Game {
	return game.NewGame(players, 100)
}

func BuildSimulator(game *game.Game) *simulator.Simulator {
	return simulator.NewSimulator(10, game)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

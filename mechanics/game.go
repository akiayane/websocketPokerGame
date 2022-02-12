package mechanics

import "fmt"

type Player struct {
	Id    int
	Cards []string
}

type Game struct {
	Players []Player
	Commons []string
}

type Winner struct {
	Player           Player
	Analyze          string
	HighestCardValue int
	Rank             int
}

func New(players []Player) Game {
	game := Game{
		Players: players,
	}

	distributeCards(&game)

	return game
}

func GetWinner(game Game) Winner {
	currentWinner := Winner{
		HighestCardValue: -1,
		Rank:             0,
		Analyze:          "invalid",
	}
	for i, _ := range game.Players {
		analyze, currentHighestCard, currentRank := AnalyzeHand(append(game.Players[i].Cards, game.Commons...))
		fmt.Println(analyze, currentHighestCard, currentRank)
		if currentRank > currentWinner.Rank {
			currentWinner.Rank = currentRank
			currentWinner.HighestCardValue = currentHighestCard
			currentWinner.Analyze = analyze
			currentWinner.Player = game.Players[i]
		} else if currentRank == currentWinner.Rank {
			if currentHighestCard > currentWinner.HighestCardValue {
				currentWinner.Rank = currentRank
				currentWinner.HighestCardValue = currentHighestCard
				currentWinner.Analyze = analyze
				currentWinner.Player = game.Players[i]
			}
		}
	}

	return currentWinner
}

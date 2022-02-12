package mechanics

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

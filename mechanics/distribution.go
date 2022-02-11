package mechanics

import (
	"math/rand"
	"time"
)

func getAllCards() []string {

	return []string{
		"2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s", "ts", "js", "qs", "ks", "as",
		"2h", "3h", "4h", "5h", "6h", "7h", "8h", "9h", "th", "jh", "qh", "kh", "ah",
		"2d", "3d", "4d", "5d", "6d", "7d", "8d", "9d", "td", "jd", "qd", "kd", "ad",
		"2c", "3c", "4c", "5c", "6c", "7c", "8c", "9c", "tc", "jc", "qc", "kc", "ac",
	}

}

func randomIndex() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 51
	return rand.Intn(max-min+1) + min
}

func distributeCards(game *Game) {
	//create map to check the card on being already taken or not
	checkmap := make(map[int]bool)
	//counter to be sure exactly 2 cards was handed out to commons
	counter := 0
	for {
		//leave loop if we already handed out 2 cards
		if counter == 3 {
			break
		}

		index := randomIndex()
		if !checkmap[index] {

			//append card to commons list
			game.Commons = append(game.Commons, getAllCards()[index])

			//increase counter so that our loop can be finished
			counter++
			checkmap[index] = true
		}

	}

	//then we handing out 3 cards for all of the players in the game, here we using same checkmap to be sure that players and commons do not have duplicates

	for i, _ := range game.Players {
		counter = 0
		for {
			//leave loop if we already handed out 3 cards
			if counter == 2 {
				break
			}

			index := randomIndex()

			if !checkmap[index] {

				//append card to players card list
				game.Players[i].Cards = append(game.Players[i].Cards, getAllCards()[index])

				//increase counter so that our loop can be finished
				counter++
				checkmap[index] = true
			}

		}
	}
}

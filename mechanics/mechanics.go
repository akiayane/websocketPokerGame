package mechanics

import (
	"sort"
	"strings"
)

type card struct {
	face byte
	suit byte
}

const faces = "23456789tjqka"
const suits = "shdc"

func isStraight(cards []card) bool {
	sorted := make([]card, 5)
	copy(sorted, cards)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].face < sorted[j].face
	})
	if sorted[0].face+4 == sorted[4].face {
		return true
	}
	if sorted[4].face == 14 && sorted[0].face == 2 && sorted[3].face == 5 {
		return true
	}
	return false
}

func isFlush(cards []card) bool {
	suit := cards[0].suit
	for i := 1; i < 5; i++ {
		if cards[i].suit != suit {
			return false
		}
	}
	return true
}

//this function analyze player's hand and return the name of combination, the highest card in the hand and the rank of combo
//rank of combo is an integer that represents the significance of combo
func AnalyzeHand(hand []string) (string, int, int) {
	highestCardValue := 0
	splitSet := make(map[string]bool)
	var split []string
	for _, s := range hand {
		if !splitSet[s] {
			splitSet[s] = true
			split = append(split, s)
		}
	}
	if len(split) != 5 {
		return "invalid", highestCardValue, 0
	}
	var cards []card

	for _, s := range split {
		if len(s) != 2 {
			return "invalid", highestCardValue, 0
		}
		fIndex := strings.IndexByte(faces, s[0])
		if fIndex == -1 {
			return "invalid", highestCardValue, 0
		}
		sIndex := strings.IndexByte(suits, s[1])
		if sIndex == -1 {
			return "invalid", highestCardValue, 0
		}
		if fIndex+2 > highestCardValue {
			highestCardValue = fIndex + 2
		}
		cards = append(cards, card{byte(fIndex + 2), s[1]})

	}

	groups := make(map[byte][]card)
	for _, c := range cards {
		groups[c.face] = append(groups[c.face], c)
	}

	switch len(groups) {
	case 2:
		for _, group := range groups {
			if len(group) == 4 {
				return "four-of-a-kind", highestCardValue, 8
			}
		}
		return "full-house", highestCardValue, 7
	case 3:
		for _, group := range groups {
			if len(group) == 3 {
				return "three-of-a-kind", highestCardValue, 4
			}
		}
		return "two-pair", highestCardValue, 3
	case 4:
		return "one-pair", highestCardValue, 2
	default:
		flush := isFlush(cards)
		straight := isStraight(cards)
		switch {
		case flush && straight:
			return "straight-flush", highestCardValue, 9
		case flush:
			return "flush", highestCardValue, 6
		case straight:
			return "straight", highestCardValue, 5
		default:
			return "high-card", highestCardValue, 1
		}
	}
}

package blackjack

import "fmt"

func main() {

	value := ParseCard("ace")
	fmt.Println(value)

	fmt.Println(FirstTurn("ace", "ace", "jack"))   // == "P"
	fmt.Println(FirstTurn("ace", "king", "ace"))   // == "S"
	fmt.Println(FirstTurn("five", "queen", "ace")) // == "H"
}

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "jack", "queen", "king", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	newCard1 := ParseCard(card1)
	newCard2 := ParseCard(card2)
	newDealerCard := ParseCard(dealerCard)
	cardScore := newCard1 + newCard2

	switch {
	case newCard1 == 11 && newCard2 == 11:
		return "P"
	case cardScore == 21 && (newDealerCard != 11 && newDealerCard != 10):
		return "W"
	case (cardScore == 21) && (newDealerCard == 11 || newDealerCard == 10):
		return "S"
	case (cardScore >= 17 && cardScore <= 20):
		return "S"
	case (cardScore >= 12 && cardScore <= 16) && (newDealerCard < 7):
		return "S"
	case (cardScore >= 12 && cardScore <= 16) && (newDealerCard >= 7):
		return "H"
	case (cardScore <= 11):
		return "H"
	}
	return "H"
}

//FirstTurn("ace", "ace", "jack") == "P"
/*

Stand (S)
Hit (H) - Esperar
Split (P)
Automatically win (W) Ganar

	scoreCards := newCard1 + newCard2
	switch scoreCards {
	case scoreCards == 22:
		return "asa"
	default:
		return "dd"
	}
*/

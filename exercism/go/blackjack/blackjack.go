package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

cardValues := map[string]int{
	"ace":	11,
	"two":		2,
	"three":	3,
	"four":		4,
	"five":		5,
	"six":		6,
	"seven":	7,
	"eight":	8,
	"nine":		9,
	"ten":	10,
	"jack": 10,
	"queen": 10,
	"king":		10,
}

value, ok := cardValues[card]

if ok {
	return value
} else {
	return 0
}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

if (card1 == "ace" && card2 == "ace") {
	return "P"
}

if (ParseCard(card1) + ParseCard(card2) == 21) {
	switch dealerCard {
		case "ace", "ten", "jack", "queen", "king": 
			return "S"
		default: 
			return "W"
	}
}

switch ParseCard(card1) + ParseCard(card2) {
	case 17, 18, 19, 20:
		return "S"
	case 12, 13, 14, 15, 16:
		switch ParseCard(dealerCard) {
			case 7, 8, 9, 10, 11:
				return "H"
			default:
				return "S"
	}
	default:
		return "H"
}

}

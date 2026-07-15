package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Game Flow:
/*
	The game starts

	First the initial cards are dealt, two to the player, and two to the dealer

	// Check if the player or dealer has black jack

	Then the player choose between 'S'tand or 'H'it

	If the player chose 'S'tand, the players turn is over and the dealers turn starts

	If the player chose 'H'it, the player gets another card and gets to choose again

	// Check if the player busts

	When the player eventually busts or stands, it's the dealers turn

	The dealer draws cards until he either gets 17 or above, or until he busts

	// Check if the dealer busts, if not; calculate the hand values and see whos is higher

	Then the game ends
*/

/*
	Suit: Hearts, Diamonds, Spades, Clubs
	Numbers: 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace
	Value: -1, 0, 1
*/
type Card struct {
	suit string
	number string
	value int
}



var playerHand []Card
var playerHandValue int

var dealerPreGameHand []Card
var dealerHand []Card
var dealerHandValue int

var userChoice string

var cards []Card
var playDeck []Card
var testDeck []Card

var playedCard Card

var suits = []string{ "Hearts", "Diamonds", "Spades", "Clubs" }
var numbers = []string{ "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace" }
//	Här är frågan hur man ska göra med Acet, för man får ju kolla om summan av alla cardValues går över 21, sen isåfall kolla om det finns ett ace med och då ta -10 på summan 
var cardValues = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "Jack": 10, "Queen": 10, "King": 10, "Ace": 11}
var suitSymbols = map[string]string{"Hearts": "♥", "Diamonds": "♦", "Spades": "♠", "Clubs": "♣"}

var uI string
var pressContinue string

var endGame bool



// Function to shuffle the deck randomly
func shuffleDeck(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func dealBeginnerCards(playDeck []Card) (playerHandUpdated []Card, dealerHandUpdated []Card){
	for i := 0; i <= 3; i++ {
		if i % 2 == 0 {
			playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "H", playerHand, dealerHand)
			// playerHandValue = calculateValue(playerHand)

		} else if i % 2 == 1 {
			playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "S", playerHand, dealerHand)
			// dealerHandValue = calculateValue(dealerHand)
		}
	}
	return playerHand, dealerHand
}

/*
1. Does the player have a natural blackjack?
	Yes → Continue to step 2.
	No → Continue to step 3.
2. Does the dealer also have a natural blackjack?
	Yes → Push.
	No → Player wins with Blackjack (usually pays 3:2).
3. Does the dealer have a natural blackjack?
	Yes → Dealer wins.
	No → Continue.
4. Did the player bust (over 21)?
	Yes → Player loses.
	No → Continue.
5. Dealer plays according to the house rules.
6. Did the dealer bust?
	Yes → Player wins.
	No → Continue.
7. Compare hand values.
	Player total > Dealer total → Player wins.
	Dealer total > Player total → Dealer wins.
	Equal totals → Push.
*/

func printPlayerHand(playerHand []Card) {
	// helper to format any hand using suit symbols and card numbers
	formatHand := func(hand []Card) string {
		s := ""
		for i, c := range hand {
			if i > 0 {
				s += " | "
			}
			sym := suitSymbols[c.suit]
			s += fmt.Sprintf("%s%s", sym, c.number)
		}
		return s
	}

	fmt.Printf("Player:\nHand: %s\nValue: %v\n\n\n", formatHand(playerHand), calculateValue(playerHand))
}

// Print both player and dealer hands (works for any number of cards)
func printHands(playerHand []Card, dealerHand []Card) {
	formatHand := func(hand []Card) string {
		s := ""
		for i, c := range hand {
			if i > 0 {
				s += " | "
			}
			sym := suitSymbols[c.suit]
			s += fmt.Sprintf("%s%s", sym, c.number)
		}
		return s
	}

	fmt.Printf("Player:\nHand: %s\nValue: %v\n\n\n", formatHand(playerHand), calculateValue(playerHand))
	fmt.Printf("Dealer:\nHand: %s\nValue: %v\n\n\n", formatHand(dealerHand), calculateValue(dealerHand))
}

// Check before the turn starts if any of the players has Black Jack
func initialDealCheck(playerHand []Card, dealerhand []Card) bool {
	if (calculateValue(playerHand) == 21 && (playerHand[0].number == "Ace" || playerHand[1].number == "Ace")) && (calculateValue(dealerHand) == 21 && (dealerHand[0].number == "Ace" || dealerHand[1].number == "Ace")) {
		printHands(playerHand, dealerHand)

		fmt.Println("Push!")
		return true
	} else if calculateValue(playerHand) == 21 && (playerHand[0].number == "Ace" || playerHand[1].number == "Ace") {
		printHands(playerHand, dealerHand)

		fmt.Println("Player has Black Jack! You win!")
		return true
	} else if calculateValue(dealerHand) == 21 && (dealerHand[0].number == "Ace" || dealerHand[1].number == "Ace") {
		printHands(playerHand, dealerHand)

		fmt.Println("Dealer has Black Jack! You loose!")
		return true
	} else {
		return false
	}
}

func playerTurnCheck(playerHandValue int) bool {
	if playerHandValue > 21 {
		return true
	}
	return false
}

func dealerTurnCheck(playerHandValue int, dealerHandValue int) bool {
	if dealerHandValue > 21 {
		fmt.Println("Dealer Bust! You win!")
		return true
	} else if dealerHandValue >= 17 {
		if dealerHandValue > playerHandValue {
			fmt.Println("Dealer won!")
			return true
		} else if dealerHandValue < playerHandValue {
			fmt.Println("Player won!")
			return true
		} else if dealerHandValue == playerHandValue {
			fmt.Println("Push!")
			return true
		}
	}
	return false
}



// Game Flow:
/*
	Then the player choose between 'S'tand or 'H'it

	If the player chose 'S'tand, the players turn is over and the dealers turn starts

	If the player chose 'H'it, the player gets another card

	// Check if the player busts

	If nothing happened, the player gets to choose again

	When the player eventually busts or stands, it's the dealers turn

	The dealer draws cards until he either gets 17 or above, or until he busts

	// Check if the dealer busts, if not; calculate the hand values and see whos is higher

	Then the game ends
*/

func dealCard() {
	// Asks the user what to do...
	fmt.Printf("( [H]it | [S]tand )")
	fmt.Scanln(&uI)

	// While the player Hit
	for {
		if uI == "H" {
			// Play a card, it gets removed from the playDeck and updated into the players and dealers hands
			playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, uI, playerHand, dealerHand)

			// Then calculates the values of both hands
			playerHandValue = calculateValue(playerHand)
			dealerHandValue = calculateValue(dealerHand)

			

			// Check if the game ended
			endGame = playerTurnCheck(playerHandValue)

			if endGame {
				printHands(playerHand, dealerHand)

				fmt.Println("Player Bust! You loose!")


				return
			}

			printPlayerHand(playerHand)

			// Prints only the first card and that cards value
			fmt.Printf("Dealer:\nHand: %v%d | ??\nValue: %v\n\n\n", suitSymbols[dealerHand[0].suit], cardValues[dealerHand[0].number], cardValues[dealerHand[0].number])
			

			// Asks the user what to do...
			fmt.Printf("( [H]it | [S]tand )")
			fmt.Scanln(&uI)

			
		// If the player stands, the dealers turn starts
		} else if uI == "S" {
			printHands(playerHand, dealerHand)

				
			// Checks if the dealer has bust or gone over 17
			endGame = dealerTurnCheck(playerHandValue, dealerHandValue)

			// If so; end the game
			if endGame {
				break
			// If not; dealer gets another card
			} else {
				// Then we wait for the player to be done reading
				fmt.Println("Press <Enter> to continue...")
				fmt.Scanln(&pressContinue)

				playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "S", playerHand, dealerHand)

				// Then calculates the values of both hands
				dealerHandValue = calculateValue(dealerHand)
			}
		}	
	}
}


// Function to calculate the value of the given hand, and if the score is over 21 and the hand has an Ace, make the Ace a 1 instead of 11
func calculateValue(hand []Card) int{
	handValue := 0
	aces := 0

	// For each card in the hand
	for _, card := range hand {
		// If the card is an Ace, increase aces and add 11 to the handValue
		if card.number == "Ace" {
			handValue += 11
			aces++
		} else {
			handValue += cardValues[card.number]
		}
	}

	for handValue > 21 && aces > 0 {
		handValue -= 10
		aces--
	}

	return handValue
}

// A function that "plays the card", removes it from the playDeck, puts it in either the playerHand or dealerHand, and updates the player- and dealerHand
func playCard(playDeck []Card, userChoice string , playerHand []Card, dealerHand []Card) (playedDeck []Card, playedCard Card, playerHandUpdated []Card, dealerHandUpdated []Card) {
	// Sets the playedCard to the first card in the deck
	playedCard = playDeck[0]

	// Puts the playedCard in either the players or the dealers hand
	if userChoice == "H" {
		playerHand = append(playerHand, playedCard)
	} else if userChoice == "S" {
		dealerHand = append(dealerHand, playedCard)
	}


	// Returns the played card, the playDeck with the played card removed, the players hand and the dealers hand
	return playDeck[1:], playedCard, playerHand, dealerHand
}

func main() {
	// Make the cards list with 52 Card structs
	cards = make([]Card, 52)

	// Give all the 52 cards in the "cards" list their unique card identity
	// For each suit
	for suitIndex, suitData := range suits {
		// For each number in each suit
		for numberIndex, numberData := range numbers {
			// i is the index, which gets calculated based on the suit and number indexes
			i := suitIndex * len(numbers) + numberIndex


			// Gives each card a value, for counting the cards, this is based on the numberIndex
			// value: 1
			if numberIndex >= 0 && numberIndex <= 4 {
				cards[i] = Card{
					suit: suitData,
					number: numberData,
					value: 1,
				}
			// value: 0
			} else if numberIndex >= 5 && numberIndex <= 7 {
				cards[i] = Card{
					suit: suitData,
					number: numberData,
					value: 0,
				}
			// value: -1
			} else {
				cards[i] = Card{
					suit: suitData,
					number: numberData,
					value: -1,
				}
			}
		}
	}
	// Creates the playDeck with 52 Card structs
	playDeck = make([]Card, len(cards))
	// testDeck = make([]Card, len(cards))

	// Shuffels the deck
	shuffleDeck(cards)

	// Copies the cards deck to the playDeck
	copy(playDeck, cards)

	// Deal the beginning cards
	playerHand, dealerHand = dealBeginnerCards(playDeck)


	
	// Calculate each hands value
	playerHandValue = calculateValue(playerHand)
	dealerHandValue = calculateValue(dealerHand)
	

	if initialDealCheck(playerHand, dealerHand) {
		return
	}

	
	// Prints the hand and value for the player hand
	printPlayerHand(playerHand)
	// fmt.Printf("Player:\nHand: %v%d | %v%d\nValue: %v\n\n\n", suitSymbols[playerHand[0].suit], cardValues[playerHand[0].number], suitSymbols[playerHand[1].suit], cardValues[playerHand[1].number], playerHandValue)

	// Prints only the first card and that cards value
	fmt.Printf("Dealer:\nHand: %v%d | ??\nValue: %v\n\n\n", suitSymbols[dealerHand[0].suit], cardValues[dealerHand[0].number], cardValues[dealerHand[0].number])
	
	

	dealCard()
}
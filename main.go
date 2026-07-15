package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
			calculateValue(playedCard, "H", playerHand, dealerHand)

		} else if i % 2 == 1 {
			playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "S", playerHand, dealerHand)
			calculateValue(playedCard, "S", playerHand, dealerHand)
			/*
			playDeck = playDeck[1:]
			dealerHand = append(dealerHand, playDeck[0])
			dealerHandValue += cardValues[playDeck[0].number]
			*/
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
func checkWin(playerHandValue int, dealerHandValue int) (endGame bool){
	if playerHandValue == 21 {
		if dealerHandValue == 21 {
			fmt.Println("Push!")
			return true
		} else {
			fmt.Println("You have a Black Jack! You win!")
			return true
		}
	} else {
		if dealerHandValue == 21 {
			fmt.Println("Dealer has Black Jack! You lost!")
			return true
		} else {
			if playerHandValue > 21 {
				fmt.Println("You Bust!")
				return true
			} else {
				if dealerHandValue > 21 {
					fmt.Println("Dealer Bust! You win!")
					return true
				} else {
					if playerHandValue > dealerHandValue {
						fmt.Println("You win!")
						return true
					} else if playerHandValue < dealerHandValue {
						fmt.Println("You lost!")
						return true
					} else if playerHandValue == dealerHandValue {
						fmt.Println("Push!")
						return true
					} else {
						return false
					}
				}
			}
		}
	}
}


func dealCard() {
	// Asks the user what to do...
	fmt.Printf("( [H]it | [S]tand )")
	fmt.Scanln(&uI)

	// Play loop
	for _ = range playDeck {
		// If the player hits
		if uI == "H" {
			// Play a card, it gets removed from the playDeck and updated into the players and dealers hands
			playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, uI, playerHand, dealerHand)

			// Then calculates the values of both hands
			playerHandValue, dealerHandValue = calculateValue(playedCard, uI, playerHand, dealerHand)

			// Check if the game ended
			endGame = checkWin(playerHandValue, dealerHandValue)

			// Prints the hand and value for the player hand
			fmt.Printf("Player:\nHand: %v\nValue: %v\n\n\n", playerHand, playerHandValue)

			// If the game is over, print the dealers whole hand
			if endGame {
				fmt.Printf("Dealer:\nHand: %v\nValue: %v\n\n\n", dealerHand, dealerHandValue)
				break
			} else {
				// Prints only the first card and that cards value
				fmt.Printf("Dealer:\nHand: %v\nValue: %v\n\n\n", dealerHand[0], cardValues[dealerHand[0].number])
				
				// Asks the user what to do...
				fmt.Printf("( [H]it | [S]tand )")
				fmt.Scanln(&uI)
			}

			/*
			if endGame {
				break
			} else {
				playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "D", playerHand, dealerHand)

				playerHandValue, dealerHandValue = calculateValue(playedCard, "D", playerHand, dealerHand)
			}
				*/
		// If the player stands
		} else if uI == "S" {

			fmt.Printf("Player:\nHand: %v\nValue: %v\n\n\n", playerHand, playerHandValue)
			fmt.Printf("Dealer:\nHand: %v\nValue: %v\n\n\n", dealerHand, dealerHandValue)

			// Then we wait for the player to be done reading
			// fmt.Println("Press <Enter> to continue...")
			// fmt.Scanln(&pressContinue)
			
			endGame = checkWin(playerHandValue, dealerHandValue)

			if endGame {
				break
			} else {
				playDeck, playedCard, playerHand, dealerHand = playCard(playDeck, "D", playerHand, dealerHand)

				// Then calculates the values of both hands
				playerHandValue, dealerHandValue = calculateValue(playedCard, "D", playerHand, dealerHand)
			}
		}
	}
}


func calculateValue(playedCard Card, userChoice string, playerHand []Card, dealerHand []Card) (playerCardValueUpdated int, dealerCardValueUpdated int){

	if userChoice == "H" {
		playerHandValue += cardValues[playedCard.number]
	} else if userChoice == "S" {
		dealerHandValue += cardValues[playedCard.number]
	}

	return playerHandValue, dealerHandValue
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
	
	fmt.Println("\n\n\nMåste skapa ett smartare system för att ta reda på när någon vinner...\nSå efter varje gång man antingen Hitar eller Standar så kollar den om någon har vunnit\n\n\n")
	fmt.Println("Måste också dela in dealCard() i mindre delar för att göra det till ett lättare och smartare system...\n\n\n")

	fmt.Println("\n\n\nSå fort man har dragit ett kort, så ska man kolla om nån har funnit. Så gör en ny funktion som heter checkWin() som anropas direkt efter att ett kort har dragits. Den kan t.o.m anropas i playCard() funktionen för att inte glömma/tappas bort")
	

	fmt.Println("LÄS IGENOM ALLA FUNKTION OCH TÄNK, TÄNK OCH TÄNK PÅ VAD SOOM SKA GÖRAS, SKRIV OM HELA PROGRAMMET OM DET BEHÖVS...")

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

	playerHand, dealerHand = dealBeginnerCards(playDeck)

	fmt.Printf("Player:\nHand: %v\nValue: %v\n\n\n", playerHand, playerHandValue)
	fmt.Printf("Dealer:\nHand: %v\nValue: %v\n\n\n", dealerHand[0], cardValues[dealerHand[0].number])
	
	

	dealCard()
}
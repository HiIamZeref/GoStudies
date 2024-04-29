package main

import "fmt"

func main() {
	var card string = "Ace of Spades" 
	//  
	// Variable declaration (var <variable_name> <variable_type> = <value>)
	// Go is a statically typed language. The type of the variable is known at the compile time.
	// The type of the variable is required to be mentioned.
	// The value of the variable can be changed later.
	card = "Five of Diamonds"
	// Variable assignment (<variable_name> = <value>)
	anotherCard := "King of Hearts"
	// Short hand declaration (var <variable_name> = <value>)
	// The type of the variable is inferred from the value

	fmt.Println("This is my card: ",card)
	fmt.Println("This is another card: ",anotherCard)

	// Basic Go types:
	var myBool bool = true // Boolean
	var myInt int = 10 // Integer
	var myFloat float64 = 3.14 // Floating point number
	var myString string = "Hello, World!" // String
	fmt.Println(myBool, myInt, myFloat, myString)

	// Getting card from the function
	// fmt.Println(newCard())
	// gottenCard := newCard()
	// fmt.Println(gottenCard)
	
	// Arrays and Slices
	// cards := []string{"Ace of Diamonds", newCard()} // Slice
	cards := newDeck() // new type: deck
	staticCards := [2]string{"Ace of Diamonds", newCard()} // Array	
	fmt.Println(cards, staticCards)

	cards = append(cards, "Six of Spades") // Append to the slice
	fmt.Println(cards)

	// Iterating over the slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining Deck: ")
	remainingDeck.print()

	// cards.saveToFile("my_cards")

	// newCards := newDeckFromFile("my_cards")
	// fmt.Println("New Cards: ")
	// newCards.print()
	fmt.Println("Original Cards: ")
	cards.print()
	cards.shuffle()
	fmt.Println("Shuffled Cards: ")
	cards.print()
}

func newCard () string {
	return "Five of Diamonds"
}
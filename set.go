package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// card represents a unique combination of the four attributes.
type card [4]int

// Attribute indices
const (
	color = iota
	shape
	fill
	number
)

// Color
const (
	red = iota
	green
	purple
)

// Shape
const (
	diamond = iota
	squiggle
	oval
)

// Fill
const (
	solid = iota
	empty
	striped
)

// Number
const (
	one = iota
	two
	three
)

// isSet takes in three cards, returning `true` if they
// compose a set and `false` if not.
func isSet(a, b, c card) bool {
	// UNIMPLEMENTED
	return true
}

// generateDeck returns a new, unshuffled slice of cards.
func generateDeck() []card {
	var buffer [4]int

	deck := make([]card, 81)

	for i := range deck {
		deck[i] = buffer
		// Inefficiency: this inner loop runs even when i = len(deck)-1
		// and calculating the next buffer state is unnecessary
		for j, val := range buffer {
			if val == 2 {
				buffer[j] = 0
				continue
			}
			if val < 2 {
				buffer[j] = val + 1
				break
			}
		}
	}

	return deck
}

// shuffleDeck performs an in-place shuffle upon a slice of cards.
func shuffleDeck(deck []card) {
	// UNIMPLEMENTED
}

// board represents a group of dealt cards (the table) and
// an associated deck from which to deal more cards.
type board struct {
	table []card
	deck  []card
}

// canDeal describes whether more cards can be dealt from
// a board's source deck. If canDeal -> false and table -> empty,
// the game is over.
func (b *board) canDeal() bool {
	return len(b.deck) > 0
}

// dealThree slices off three cards from the board's source deck
// and appends them to the board's table.
func (b *board) dealThree() {
	b.table = append(b.table, b.deck[0:3]...)
	b.deck = b.deck[3:]
}

// dealTwelve slices off twelve cards from the board's source deck
// and appends them to the board's table. Only occurs at game start.
func (b *board) dealTwelve() {
	b.table = append(b.table, b.deck[0:12]...)
	b.deck = b.deck[12:]
}

// startGame receives a deck and returns the address of a new board.
// The board retains a reference to its source deck for further dealing.
func startGame(deck []card) *board {
	// UNIMPLEMENTED
	return &board{}
}

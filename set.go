package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	game := newGame()
	var done bool

	for !done {
		done = game.play()
	}

	for i, set := range game.sets {
		fmt.Printf("Set %2d:\n\t%v\n\t%v\n\t%v\n", i+1, set[0], set[1], set[2])
	}
}

// card represents a unique combination of the four attributes.
type card [4]int

// Attribute index constants
const (
	color = iota
	shape
	fill
	number
)

// Number
var numbers = map[int]string{
	0: "one",
	1: "two",
	2: "three",
}

// Fill
var fills = map[int]string{
	0: "solid",
	1: "empty",
	2: "striped",
}

// Color
var colors = map[int]string{
	0: "red",
	1: "green",
	2: "purple",
}

// Shape
var shapes = map[int]string{
	0: "diamond",
	1: "squiggle",
	2: "oval",
}

// card.String returns a stringifed representation of the card's attributes.
func (c card) String() string {
	attributes := []string{
		numbers[c[number]],
		fills[c[fill]],
		colors[c[color]],
		shapes[c[shape]],
	}
	if c[number] > 0 {
		attributes[3] = attributes[3] + "s"
	}
	return strings.Join(attributes, " ")
}

// isSet takes in three cards, returning `true` if they
// compose a set and `false` if not.
//
// isSet is not actually used in any of the other set-finding
// functions; instead, the rest of the program uses complement-checking
// to identify sets.
func isSet(a, b, c card) bool {
	// Because we've modeled attributes as integers 0 <= n <= 2,
	// we can test validity by taking a sum of each card's values
	// for a given attribute. There are four valid sums --
	// three for homogenous attributes:
	// 0 + 0 + 0 = 0,
	// 1 + 1 + 1 = 3,
	// 2 + 2 + 2 = 6.
	// and one for heterogeneous attributes:
	// 0 + 1 + 2 = 3.
	// A modulo by three can check for all of these at once.
	for i := range a {
		if (a[i]+b[i]+c[i])%3 != 0 {
			return false
		}
	}
	return true
}

// complement takes two cards and returns a representation of the
// card that, if found, would complete a set with the two given cards.
func complement(a, b card) card {
	var c card

	for i := range c {
		if a[i] == b[i] {
			// If attribute `i` is homogeneous between a and b,
			// c should be the same.
			c[i] = a[i]
		} else {
			// If attribute `i` is heterogeneous, c must be complement a and b.
			// Attributes (as integers in range [0..2]) are represented by two bits.
			// By using 3 (0b11) as a mask for a XOR against the sum of a[i] and b[i],
			// we can calculate the necessary complement.
			// For example:
			// 0 + 1 = 1 (0b01). 0b01 ^ 0b11 = 0b10 = 2, the correct complement to 0 and 1.
			// 1 + 2 = 3 (0b11). 0b11 ^ 0b11 = 0b00 = 0, etc.
			c[i] = (a[i] + b[i]) ^ 3
		}
	}

	return c
}

// generateDeck returns a new, unshuffled slice of cards.
// It modifies a single integer array and repeatedly copies it
// into the deck to create 81 unique cards. The process of modifying
// the buffer array is conceptually similar to counting from 0 to 80
// in big-endian base 3, with each index as a digit of the number.
func generateDeck() []card {
	var buffer [4]int

	deck := make([]card, 81)

	for i := range deck {
		deck[i] = buffer
		// Inefficiency: this inner loop runs even when i = len(deck)-1,
		// at which point calculating the next buffer state is unnecessary
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
	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		randIx := rand.Intn(i + 1)
		if randIx != i {
			temp := deck[randIx]
			deck[randIx] = deck[i]
			deck[i] = temp
		}
	}
}

// board represents a group of dealt cards (the table) and
// an associated deck from which to deal more cards.
type board struct {
	table []card
	deck  []card
	sets  [][3]card
}

// canDeal describes whether more cards can be dealt from
// a board's source deck. If canDeal -> false and table -> empty,
// the game is over.
func (b *board) canDeal() bool {
	return len(b.deck) >= 3
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

// findSet searches the board's table for a valid set.
// If a valid set is found, the cards are returned and
// `found` is true; if there's no valid set, `set` is nil
// and `found` is false.
//
// For each card in `set`, there's both the card value as well
// as the index in the table slice at which it was found.
// This enables later removal of these cards from the table.
func (b *board) findSet() (set [3]card, found bool) {
	// First, we create a map of cards to indices by looping through
	// the table and inserting each card. This helps amortize the work
	// of checking permutations of sets, since for any pair of cards
	// we can calculate which third card would complete their set.
	// With a map of all the cards on the table, we can do this third
	// check in O(1) time, instead of a third loop of O(n) complexity.
	mem := make(map[card]int)

	for i, card := range b.table {
		mem[card] = i
	}

	// We label this outer loop so we can break the entire search
	// operation once a set has been identified in the inner loop.
	//
	// An alternative would be to return immediately upon identifying
	// the set, but using this break allows us to use only a single
	// return statement, which is part structural, part stylistic decision.
outer:
	for i := range b.table {
		for j := i + 1; j < len(b.table); j++ {
			// Here, for each pair of cards as we iterate through the table,
			// we calculate which third card would create a complete set if present.
			comp := complement(b.table[i], b.table[j])

			// Next, we check to see if that third card is on the table.
			if _, ok := mem[comp]; ok {
				// If it is, we mark our `found` return value as true, and we instantiate
				// and assign a results array containing the cards of the set we've found.
				found = true
				set = [3]card{b.table[i], b.table[j], comp}
				// Here, we break the outer loop, since a set has been found and there's
				// no point in continuing to look.
				break outer
			}
		}
	}
	// Because our return values are named, we use a bare return here.
	// If we found a set, the values have been properly assigned.
	// If we haven't, we return nil/false, which is useful.
	return
}

// clearSet takes a set of cards as input and replaces the board's
// table with a slice omitting the cards of the set.
func (b *board) clearSet(set [3]card) {
	// Here we create the slice that will become the new table.
	// We want to shift cards around to new indices to fill the gaps,
	// which `append` makes convenient, but we also don't want to
	// repeatedly allocate a new underlying array when it's not
	// necessary, so we create the slice with length 0 but adequate capacity.
	t := make([]card, 0, 27)

outer:
	for _, card := range b.table {
		for _, setMember := range set {
			if setMember == card {
				continue outer
			}
		}
		t = append(t, card)
	}

	b.table = t
}

// newGame creates and shuffles a new deck, associates this deck with
// a new board, deals out twelve cards to the board's table, then returns
// the address of the board.
func newGame() *board {
	deck := generateDeck()
	shuffleDeck(deck)

	b := &board{deck: deck}
	b.dealTwelve()

	return b
}

// play executes one "tick" of the game -- it attempts to find a set, and if it
// succeeds, it removes that set and appends it to a log of sets found, then deals
// out more cards if necessary. If no set is found, play will deal out three more
// cards, then return and cede control. play returns true if the game is finished
// (i.e., no more sets and no more cards to deal) and false if the game can continue.
func (b *board) play() (done bool) {
	set, found := b.findSet()

	if found {
		b.sets = append(b.sets, set)
		b.clearSet(set)
		done = false
		if b.canDeal() {
			b.dealThree()
		}
	} else {
		if b.canDeal() {
			b.dealThree()
		} else {
			done = true
		}
	}

	return
}

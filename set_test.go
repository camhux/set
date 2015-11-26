package main

import (
	"testing"
)

func TestIsSet(t *testing.T) {
	cases := [...]struct {
		set      [3]card
		expected bool
	}{
		{
			set: [3]card{
				{0, 1, 1, 0},
				{1, 1, 2, 0},
				{2, 0, 0, 0},
			},
			expected: false,
		},
		{
			set: [3]card{
				{0, 1, 1, 0},
				{1, 1, 2, 0},
				{2, 1, 0, 0},
			},
			expected: true,
		},
		{
			set: [3]card{
				{2, 2, 2, 2},
				{0, 0, 0, 0},
				{1, 1, 1, 1},
			},
			expected: true,
		},
		{
			set: [3]card{
				{2, 2, 2, 1},
				{0, 1, 0, 0},
				{1, 1, 1, 1},
			},
			expected: false,
		},
		{
			set: [3]card{
				{0, 2, 2, 1},
				{0, 0, 0, 1},
				{1, 1, 1, 1},
			},
			expected: false,
		},
	}

	for _, tCase := range cases {
		set := tCase.set

		expected := tCase.expected
		actual := isSet(set[0], set[1], set[2])

		if expected != actual {
			t.Errorf("Expected %v from `isSet`; got %v\n", expected, actual)
		}
	}
}

func TestComplement(t *testing.T) {
	cases := [...]struct {
		a, b, expected card
	}{
		{
			a:        card{0, 0, 1, 2},
			b:        card{0, 1, 1, 0},
			expected: card{0, 2, 1, 1},
		},
		{
			a:        card{2, 0, 1, 0},
			b:        card{0, 2, 1, 1},
			expected: card{1, 1, 1, 2},
		},
		{
			a:        card{0, 0, 0, 1},
			b:        card{0, 2, 1, 1},
			expected: card{0, 1, 2, 1},
		},
		{
			a:        card{1, 2, 2, 0},
			b:        card{2, 2, 1, 1},
			expected: card{0, 2, 0, 2},
		},
		{
			a:        card{1, 0, 0, 0},
			b:        card{2, 2, 1, 1},
			expected: card{0, 1, 2, 2},
		},
	}

	for _, tCase := range cases {
		expected := tCase.expected
		actual := complement(tCase.a, tCase.b)

		if expected != actual {
			t.Errorf("Expected %v from `complement`; got %v\n", expected, actual)
		}
	}
}

func TestCanDeal(t *testing.T) {
	var b board
	var expected, actual bool

	check := func(expected, actual bool) {
		if expected != actual {
			t.Errorf("Expected %v from `b.canDeal`; got %v instead", expected, actual)
		}
	}

	b = board{}
	b.deck = append(b.deck, card{}, card{}, card{})

	expected, actual = true, b.canDeal()
	check(expected, actual)

	b = board{}
	b.deck = append(b.deck, card{})

	expected, actual = false, b.canDeal()
	check(expected, actual)

}

func TestDealThree(t *testing.T) {

}

func TestDealTwelve(t *testing.T) {

}

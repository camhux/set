package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

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
	var b board
	var expected, actual []card

	check := func(expected, actual []card) {
		if len(expected) != len(actual) {
			t.Error("Mismatched length between expected table and actual")
		}

		for i := range expected {
			if expected[i] != actual[i] {
				t.Error("Mismatched cards between expected table and actual")
			}
		}
	}

	b = board{}
	b.deck = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 1},
		{0, 0, 2, 2},
	}

	b.dealThree()
	expected = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
	}
	actual = b.table
	check(expected, actual)

	b.dealThree()
	expected = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
	}
	actual = b.table
	check(expected, actual)

}

func TestDealTwelve(t *testing.T) {
	var b board
	var expected, actual []card

	check := func(expected, actual []card) {
		if len(expected) != len(actual) {
			t.Error("Mismatched length between expected table and actual")
		}

		for i := range expected {
			if expected[i] != actual[i] {
				t.Error("Mismatched cards between expected table and actual")
			}
		}
	}

	b = board{}
	b.deck = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 1},
		{0, 0, 2, 2},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 2},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{0, 1, 1, 2},
	}

	b.dealTwelve()
	expected = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 1},
		{0, 0, 2, 2},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 2},
	}
	actual = b.table
	check(expected, actual)
}

func TestFindSet(t *testing.T) {
	var b board
	var expected, actual [3]card
	var ok bool

	check := func(expected, actual [3]card) {
		if expected != actual {
			t.Errorf("Expected %v from `b.findSet`; got %v", expected, actual)
		}
	}

	b.table = []card{
		{0, 0, 0, 0},
		{0, 1, 1, 1},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
		{1, 0, 1, 0},
		{2, 0, 2, 0},
	}

	expected = [3]card{
		{0, 0, 0, 0},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
	}
	actual, _ = b.findSet()
	check(expected, actual)

	b.table = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 2},
		{1, 0, 1, 0},
		{2, 1, 1, 0},
		{1, 1, 1, 1},
	}
	expected = [3]card{}
	actual, ok = b.findSet()
	if ok {
		t.Error("Expected set to not be found")
	}
	check(expected, actual)

}

func TestClearSet(t *testing.T) {
	var b board
	var set [3]card
	var expected, actual []card

	check := func(expected, actual []card) {
		if len(expected) != len(actual) {
			t.Error("Mismatched lengths between expected table and actual")
		}

		for i := range expected {
			if expected[i] != actual[i] {
				t.Error("Mismatched cards between expected table and actual")
			}
		}
	}

	b.table = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 1},
		{0, 0, 2, 2},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 2},
	}

	set = [3]card{
		{0, 0, 1, 0},
		{0, 0, 2, 1},
		{0, 1, 0, 2},
	}

	expected = []card{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 2},
		{0, 0, 1, 1},
		{0, 0, 1, 2},
		{0, 0, 2, 0},
		{0, 0, 2, 2},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
	}

	b.clearSet(set)
	actual = b.table

	check(expected, actual)

}

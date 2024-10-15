package main

import (
	"fmt"
	"slices"
	"testing"
)

type TestTable struct {
	ExpectedStates States
	Question       string
}

func TestTruthTable(t *testing.T) {
	tests := []TestTable{
		{
			ExpectedStates: States{
				State("a"):         []Binary{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				State("b"):         []Binary{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				State("c"):         []Binary{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
				State("d"):         []Binary{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
				State("a'"):        []Binary{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				State("b'"):        []Binary{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
				State("c'"):        []Binary{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				State("d'"):        []Binary{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
				State("a.b.c"):     []Binary{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
				State("c.d"):       []Binary{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1},
				State("a.b.c+c.d"): []Binary{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
			},
			Question: "abc + cd",
		}, {
			ExpectedStates: States{
				State("a"):          []Binary{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				State("b"):          []Binary{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				State("c"):          []Binary{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
				State("d"):          []Binary{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
				State("a'"):         []Binary{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				State("b'"):         []Binary{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
				State("c'"):         []Binary{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				State("d'"):         []Binary{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
				State("a'.b.c"):     []Binary{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				State("c.d"):        []Binary{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1},
				State("a'.b.c+c.d"): []Binary{0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1},
			},
			Question: "a'.b.c + c.d",
		},
		{
			ExpectedStates: States{
				State("a"):             []Binary{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				State("b"):             []Binary{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				State("c"):             []Binary{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
				State("d"):             []Binary{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
				State("a'"):            []Binary{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				State("b'"):            []Binary{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
				State("c'"):            []Binary{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				State("d'"):            []Binary{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
				State("(a'+b+c)"):      []Binary{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1},
				State("(c+d)"):         []Binary{0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1},
				State("(a'+b+c)(c+d)"): []Binary{0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1},
			},
			Question: "(a' + b + c)(c + d)",
		},
		{
			ExpectedStates: States{
				State("a"):  []Binary{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
				State("a'"): []Binary{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},

				State("b"):  []Binary{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1},
				State("b'"): []Binary{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0},

				State("c"):  []Binary{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
				State("c'"): []Binary{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0},

				State("d"):  []Binary{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
				State("d'"): []Binary{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},

				State("(a+b)"):                  []Binary{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				State("(d+b)"):                  []Binary{0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1},
				State("(c+b)"):                  []Binary{0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1},
				State("(a'+b')"):                []Binary{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
				State("(a+b)(d+b)(c+b)(a'+b')"): []Binary{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0},
			},
			Question: "(a+b)(d+b)(c+b)(a'+b')",
		},
	}
	for _, test := range tests {
		tokens, stateNames, _ := GenerateTokensAndStates(test.Question)
		terms, isPos := ParseTerms(&tokens)
		states, termStrings := CalculateTermBinaries(
			terms,
			isPos,
			PopulatesStateBins(tokens, stateNames),
		)
		states = CalculateFinalTable(&termStrings, isPos, states)
		if !mapEqual(states, test.ExpectedStates) {
			map1 := "\n"
			map2 := "\n"
			for k, v := range test.ExpectedStates {
				map1 += string(k) + " " + fmt.Sprint(v) + "\n"
				map2 += string(k) + " " + fmt.Sprint(states[k]) + "\n"
				if !slices.Equal(v, states[k]) {
					t.Errorf("Expected %v\nGot %v", v, states[k])
				}
			}
			t.Errorf("Expected: %v\nGot: %v", map1, map2)
		}
	}
}

func mapEqual(m1 States, m2 States) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		v1, e1 := m1[k]
		v2, e2 := m2[k]
		if e1 != e2 {
			return false
		}
		if !slices.Equal(v1, v2) {
			return false
		}
	}
	return true
}

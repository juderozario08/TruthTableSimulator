package main

import (
	"github.com/golang-collections/collections/stack"
)

type Expression struct {
	Term []any
	Type int // Can be SOP POS or NEITHER
}

const (
	SOP = iota
	POS
	NEITHER
)

func ValidateExpressionAndParseTerms(expr string, terms *Expression) bool {
	st := stack.New()
	switch {
	case expr[0] == '(':
		st.Push(expr[0])
		str := make([]byte, 0)
		for i := 1; i < len(expr); i++ {
			switch expr[i] {
			case '(':
				st.Push(expr[i])
			case ')':
				st.Pop()
				if st.Len() == 0 {
					terms.Term = append(terms.Term, parseTerm(str))
					str = make([]byte, 0)
				}
			default:
				str = append(str, expr[i])
			}
		}
	default:
	}
	return false
}

func parseTerm(_ []byte) any {
	return nil
}

func parseSOP(_ []byte) any {
	return nil
}

func parsePOS(_ []byte) any {
	return nil

}

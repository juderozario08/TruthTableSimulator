package main

func ParseTerms(tokens *[]Token) (terms []Term, isPos bool) {
	terms = make([]Term, 0)
	pos := 0
	for i := 0; i < len(*tokens); i++ {
		switch (*tokens)[i].Type {
		case TokenBracketOpen:
			terms = append(terms, ParseSOP(tokens, &i))
		default:
			terms = append(terms, ParsePOS(tokens, &i))
			pos = 1
		}
	}
	if pos == 1 {
		return terms, true
	}
	return terms, false
}

func ParseSOP(tks *[]Token, i *int) Term {
	term := make(Term, 0)
	term = append(term, (*tks)[*i])
	*i++
	for ; (*tks)[*i].Type != TokenBracketClose; *i++ {
		term = append(term, (*tks)[*i])
	}
	term = append(term, (*tks)[*i])
	return term
}

func ParsePOS(tks *[]Token, i *int) Term {
	term := make(Term, 0)
	for ; *i < len(*tks) && (*tks)[*i].Type != TokenOr; *i++ {
		term = append(term, (*tks)[*i])
	}
	return term
}

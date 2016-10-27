package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n\n--------------\nFalses\n----------------\n")
	falses := []string{
		"1+",
		"1c",
		"1+((2*3)",
		"1+2/4*c",
		"+2",
		"()",
		"1)(()",
		"1+((*3)",
		"",
		"+",
	}
	run(falses)

	// true
	fmt.Println("\n\n--------------\nTrues\n----------------\n")
	trues := []string{
		"1+2",
		"1+(2*3)",
		"1+((2*3))",
		"1+2/4",
		"1 + 2 / 334",
		"12323 + (7892) / (334*4)",
		"(1+2)/4",
		"(12323 + (7892) / (334*4)) + 12323 + (7892) / (334*4)",
	}
	run(trues)
}

func run(cases []string) {
	for _, c := range cases {
		fmt.Println(fmt.Sprintf("%v <- %q", accept(c), c))
	}
}

func accept(expr string) bool {
	s := &scanner{runes: []rune(strings.Replace(expr, " ", "", -1))}
	return isExpr(s) && s.done()
}

// LPAREN = (
// RPAREN = )
// DIGIT = 1|2|3|4|5|6|7|8|9
// NUMBER = <DIGIT>|<DIGIT><NUMBER>
// OP = +|*|/|-
// EXPR = <NUMBER>|<NUMBER><OP><EXPR>|<LPAREN><EXPR><RPAREN>|<EXPR><OP><EXPR>

type scanner struct {
	i     int
	runes []rune
}

func (s *scanner) cur() rune {
	if s.i >= len(s.runes) {
		return rune(0)
	}
	return s.runes[s.i]
}

func (s *scanner) peekNext() rune {
	if s.i >= len(s.runes)-1 {
		return rune(0)
	}
	return s.runes[s.i+1]
}

func (s *scanner) adv() {
	s.i += 1
}

func (s *scanner) done() bool {
	return s.i == len(s.runes)
}

func isExpr(s *scanner) bool {
	if isDigit(s.cur()) {
		for isDigit(s.peekNext()) { // Loop to lop off digits
			s.adv()
		}
		s.adv()
	} else if s.cur() == '(' {
		s.adv()
		if !isExpr(s) {
			return false
		}
		if s.cur() != ')' {
			return false
		}
		s.adv()
	} else {
		return false
	}

	if isOp(s.cur()) {
		s.adv()
		return isExpr(s)
	}
	return true
}

func isDigit(r rune) bool {
	switch r {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func isOp(r rune) bool {
	switch r {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

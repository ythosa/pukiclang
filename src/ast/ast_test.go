package ast

import (
	"testing"

	"github.com/ythosa/pukiclang/src/token"
)

func TestString(t *testing.T) {
	program := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "tan9"},
					Value: "tan9",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "yaya"},
					Value: "yaya",
				},
			},
		},
	}

	if program.String() != "let tan9 = yaya;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

package parser

import (
	"reflect"
	"testing"
	"unicode"

	"github.com/fbac/calc-o-matic/pkg/grammar"
	"github.com/fbac/calc-o-matic/pkg/stack"
)

func TestParseExpr(t *testing.T) {
	type args struct {
		inputExpr string
	}
	tests := []struct {
		name    string
		args    args
		want    stack.Stack
		wantErr bool
	}{
		{
			"ParseExpr - Success",
			args{inputExpr: "2 * 2 + 2"},
			stack.Stack{
				grammar.Token{T: 1, V: "2"},
				grammar.Token{T: 0, V: "*"},
				grammar.Token{T: 1, V: "2"},
				grammar.Token{T: 0, V: "+"},
				grammar.Token{T: 1, V: "2"},
			},
			false,
		},
		{
			"ParseExpr - Fail (invalid expr)",
			args{inputExpr: "1 + 6 + a"},
			nil,
			true,
		},
		{
			"ParseExpr - Fail (empty)",
			args{inputExpr: "  "},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseExpr(tt.args.inputExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAST(t *testing.T) {
	type args struct {
		inputStack stack.Stack
	}

	validInput := stack.Stack{
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "*"},
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "+"},
		grammar.Token{T: 1, V: "2"},
	}

	invalidInput := stack.Stack{}

	tests := []struct {
		name    string
		args    args
		want    stack.Stack
		wantErr bool
	}{
		{
			"CreateAST - Success",
			args{inputStack: validInput},
			stack.Stack{
				grammar.Token{T: 1, V: "2"},
				grammar.Token{T: 1, V: "2"},
				grammar.Token{T: 0, V: "*"},
				grammar.Token{T: 1, V: "2"},
				grammar.Token{T: 0, V: "+"},
			},
			false,
		},
		{
			"CreateAST - Fail (invalid expr)",
			args{inputStack: invalidInput},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAST(tt.args.inputStack)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAST() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"isNum - Success",
			args{s: "0"},
			true,
		},
		{
			"isNum - Fail (not a number)",
			args{s: "NaN"},
			false,
		},
		{
			"isNum - Fail (empty)",
			args{s: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNum(tt.args.s); got != tt.want {
				t.Errorf("isNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"isOp - Success",
			args{s: "+"},
			true,
		},
		{
			"isOp - Fail (not an op)",
			args{s: "NaOp"},
			false,
		},
		{
			"isOp - Fail (empty)",
			args{s: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOp(tt.args.s); got != tt.want {
				t.Errorf("isOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_runeIsNumber(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"runeIsNumber - Success",
			args{r: '1'},
			true,
		},
		{
			"runeIsNumber - Fail (not an op)",
			args{r: 'X'},
			false,
		},
		{
			"runeIsNumber - Fail (empty)",
			args{r: ' '},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runeIsNumber(tt.args.r); got != tt.want {
				t.Errorf("runeIsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"isValidLength - Success",
			args{s: "abc"},
			true,
		},
		{
			"isValidLength - Fail (empty)",
			args{s: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidLength(tt.args.s); got != tt.want {
				t.Errorf("isValidLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidInput(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"isValidInput - Match valid numbers",
			args{s: "123"},
			true,
		},
		{
			"isValidInput - Match invalid chars",
			args{s: "abc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidInput(tt.args.s); got != tt.want {
				t.Errorf("isValidInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutateString(t *testing.T) {
	type args struct {
		str string
		f   stringMutator
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"mutateString - trim whitespaces",
			args{str: "1 2 3", f: func(r rune) rune {
				if unicode.IsSpace(r) {
					return -1
				}
				return r
			}},
			"123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mutateString(tt.args.str, tt.args.f); got != tt.want {
				t.Errorf("mutateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWhitespace(t *testing.T) {
	type args struct {
		r rune
	}

	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			"isWhitespace - Success",
			args{r: ' '},
			int32(-1),
		},
		{
			"isWhitespace - Fail (empty)",
			args{r: 'N'},
			int32(78),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWhitespace(tt.args.r); got != tt.want {
				t.Errorf("isWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

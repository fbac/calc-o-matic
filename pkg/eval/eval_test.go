package eval

import (
	"testing"

	"github.com/fbac/calc-o-matic/pkg/grammar"
	"github.com/fbac/calc-o-matic/pkg/stack"
)

func TestCalculate(t *testing.T) {

	astStack1 := stack.Stack{
		grammar.Token{T: 1, V: "1"},
		grammar.Token{T: 1, V: "1"},
		grammar.Token{T: 0, V: "*"},
	}

	//[{1 2} {1 2} {0 ^} {1 1} {1 4} {0 *} {0 +} {1 4} {0 -}]
	astStack2 := stack.Stack{
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "^"},
		grammar.Token{T: 1, V: "1"},
		grammar.Token{T: 1, V: "4"},
		grammar.Token{T: 0, V: "*"},
		grammar.Token{T: 0, V: "+"},
		grammar.Token{T: 1, V: "4"},
		grammar.Token{T: 0, V: "-"},
	}

	type args struct {
		inputStack stack.Stack
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Calculate '1 / 1'",
			args{inputStack: astStack1},
			1,
		},
		{
			"Calculate '2^2 + 1 * 4 -4'",
			args{inputStack: astStack2},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.inputStack); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

package stack

import (
	"reflect"
	"testing"

	"github.com/fbac/calc-o-matic/pkg/grammar"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		{
			"NewStack - Success",
			Stack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_isValid(t *testing.T) {

	validStack := Stack{
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "*"},
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "+"},
		grammar.Token{T: 1, V: "2"},
	}

	invalidStack := Stack{}

	tests := []struct {
		name string
		s    Stack
		want bool
	}{
		{
			"isValid - Success",
			validStack,
			true,
		},
		{
			"isValid - Fail",
			invalidStack,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.isValid(); got != tt.want {
				t.Errorf("Stack.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_First(t *testing.T) {

	validStack := Stack{
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "*"},
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "+"},
		grammar.Token{T: 1, V: "2"},
	}

	invalidStack := Stack{}

	tests := []struct {
		name    string
		s       Stack
		want    grammar.Token
		want1   Stack
		wantErr bool
	}{
		{
			"First - Success",
			validStack,
			grammar.Token{T: 1, V: "2"},
			validStack[1:],
			false,
		},
		{
			"First - Fail",
			invalidStack,
			grammar.Token{},
			invalidStack,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.First()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.First() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Stack.First() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {

	validStack := Stack{
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "*"},
		grammar.Token{T: 1, V: "2"},
		grammar.Token{T: 0, V: "+"},
		grammar.Token{T: 1, V: "2"},
	}

	invalidStack := Stack{}

	tests := []struct {
		name    string
		s       Stack
		want    grammar.Token
		want1   Stack
		wantErr bool
	}{
		{
			"Pop - Success",
			validStack,
			grammar.Token{T: 1, V: "2"},
			validStack[:len(validStack)-1],
			false,
		},
		{
			"Pop - Fail",
			invalidStack,
			grammar.Token{},
			invalidStack,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Stack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStack_Merge(t *testing.T) {

	// Initial stacks
	stack0 := Stack{grammar.Token{T: 1, V: "1"}}
	stack1 := Stack{grammar.Token{T: 0, V: "+"}}

	// Ordered by descending prio
	sFinal := Stack{grammar.Token{T: 0, V: "+"}, grammar.Token{T: 1, V: "1"}}

	type args struct {
		merge Stack
	}
	tests := []struct {
		name string
		orig Stack
		args args
		want Stack
	}{
		{
			"Merge - Success",
			stack0,
			args{merge: stack1},
			sFinal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.orig.Merge(tt.args.merge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Shift(t *testing.T) {

	stack0 := Stack{grammar.Token{T: 1, V: "1"}}

	// sFinal should be an empty stack
	sFinal := Stack{}

	tests := []struct {
		name  string
		s     Stack
		want  grammar.Token
		want1 Stack
	}{
		{
			"Shift - Success",
			stack0,
			grammar.Token{T: 1, V: "1"},
			sFinal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.Shift()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Shift() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Stack.Shift() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSetStackPriority(t *testing.T) {

	stack0 := Stack{grammar.Token{T: 1, V: "1"}}
	stack1 := Stack{grammar.Token{T: 0, V: "("}}
	stack2 := Stack{grammar.Token{T: 0, V: "+"}}
	stack3 := Stack{grammar.Token{T: 0, V: "*"}}
	stack4 := Stack{grammar.Token{T: 0, V: "^"}}

	type args struct {
		s Stack
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"SetStackPriority - Prio 0",
			args{s: stack0},
			0,
		},
		{
			"SetStackPriority - Prio 1",
			args{s: stack1},
			1,
		},
		{
			"SetStackPriority - Prio 2",
			args{s: stack2},
			2,
		},
		{
			"SetStackPriority - Prio 3",
			args{s: stack3},
			3,
		},
		{
			"SetStackPriority - Prio 4",
			args{s: stack4},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetStackPriority(tt.args.s); got != tt.want {
				t.Errorf("SetStackPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenPriority(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"GetTokenPriority - NUM",
			args{s: "1"},
			0,
		},
		{
			"GetTokenPriority - (",
			args{s: "("},
			1,
		},
		{
			"GetTokenPriority - +",
			args{s: "+"},
			2,
		},
		{
			"GetTokenPriority - -",
			args{s: "-"},
			2,
		},
		{
			"GetTokenPriority - /",
			args{s: "/"},
			3,
		},
		{
			"GetTokenPriority - *",
			args{s: "*"},
			3,
		},
		{
			"GetTokenPriority - ^",
			args{s: "^"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenPriority(tt.args.s); got != tt.want {
				t.Errorf("GetTokenPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrioIsBigger(t *testing.T) {
	type args struct {
		currPrio int
		lastPrio int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"PrioIsBigger - Success",
			args{
				currPrio: 1,
				lastPrio: 0,
			},
			true,
		},
		{
			"PrioIsBigger - False",
			args{
				currPrio: 0,
				lastPrio: 1,
			},
			false,
		},
		{
			"PrioIsBigger - False",
			args{
				currPrio: 0,
				lastPrio: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrioIsBigger(tt.args.currPrio, tt.args.lastPrio); got != tt.want {
				t.Errorf("PrioIsBigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

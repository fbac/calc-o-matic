package stack

import (
	"reflect"
	"testing"
	"unicode"
)

func TestNewRuneStack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want RuneStack
	}{
		{
			"NewRuneStack - Success",
			args{s: "NewRuneStack"},
			RuneStack{78, 101, 119, 82, 117, 110, 101, 83, 116, 97, 99, 107},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuneStack(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuneStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneStack_Read(t *testing.T) {

	rStackDigit := NewRuneStack("12 + 34 ABCD")

	type args struct {
		f nextRune
	}
	tests := []struct {
		name  string
		r     RuneStack
		args  args
		want  string
		want1 RuneStack
	}{
		{
			"Read - Success",
			rStackDigit,
			args{f: func(r rune) bool {
				return unicode.IsDigit(r)
			}},
			"12",
			RuneStack{32, 43, 32, 51, 52, 32, 65, 66, 67, 68},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.Read(tt.args.f)
			if got != tt.want {
				t.Errorf("RuneStack.Read() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RuneStack.Read() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRuneStack_Shift(t *testing.T) {

	rStack := RuneStack{78, 101, 119, 82, 117, 110, 101, 83, 116, 97, 99, 107}

	tests := []struct {
		name  string
		r     RuneStack
		want  rune
		want1 []rune
	}{
		{
			"Shift - Success",
			rStack,
			rStack[0],
			rStack[1:],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.Shift()
			if got != tt.want {
				t.Errorf("RuneStack.Shift() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("RuneStack.Shift() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

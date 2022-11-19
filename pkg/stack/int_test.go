package stack

import (
	"reflect"
	"testing"
)

func TestNewIntStack(t *testing.T) {
	tests := []struct {
		name string
		want IntStack
	}{
		{
			"NewIntStack - Success",
			IntStack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntStack_isValid(t *testing.T) {

	validStack := NewIntStack()
	validStack = append(validStack, 1)
	validStack = append(validStack, 2)
	validStack = append(validStack, 3)
	validStack = append(validStack, 4)

	invalidStack := NewIntStack()

	tests := []struct {
		name string
		i    IntStack
		want bool
	}{
		{
			"IntStack isValid - Success",
			validStack,
			true,
		},
		{
			"IntStack isValid - Fail",
			invalidStack,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.isValid(); got != tt.want {
				t.Errorf("IntStack.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntStack_Pop(t *testing.T) {

	validStack := NewIntStack()
	validStack = append(validStack, 1)
	validStack = append(validStack, 2)
	validStack = append(validStack, 3)
	validStack = append(validStack, 4)

	invalidStack := NewIntStack()

	tests := []struct {
		name  string
		i     IntStack
		want  int
		want1 IntStack
	}{
		{
			"IntStack Pop - Success",
			validStack,
			4,
			validStack[:len(validStack)-1],
		},
		{
			"IntStack Pop - Fail",
			invalidStack,
			0,
			invalidStack,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.i.Pop()
			if got != tt.want {
				t.Errorf("IntStack.Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("IntStack.Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIntStack_UpdateResult(t *testing.T) {

	validStack := NewIntStack()
	validStack = append(validStack, 9999)

	tests := []struct {
		name string
		i    IntStack
		want int
	}{
		{
			"UpdateResult - Success",
			validStack,
			9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.UpdateResult(); got != tt.want {
				t.Errorf("IntStack.UpdateResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

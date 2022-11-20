package eval

import (
	"testing"
)

func TestDoFunc(t *testing.T) {
	type args struct {
		op string
		x  int
		y  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Sum",
			args{
				op: "+",
				x:  1,
				y:  1,
			},
			2,
		},
		{
			"Sub",
			args{
				op: "-",
				x:  1,
				y:  1,
			},
			0,
		},
		{
			"Mul",
			args{
				op: "*",
				x:  10,
				y:  10,
			},
			100,
		},
		{
			"/",
			args{
				op: "/",
				x:  2,
				y:  4,
			},
			2,
		},
		{
			"Pow",
			args{
				op: "^",
				x:  2,
				y:  2,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoFunc(tt.args.op, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("DoFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Sum",
			args{x: 1, y: 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Sub - Positive",
			args{x: 1, y: 1},
			0,
		},
		{
			"Sub - Negative",
			args{x: 0, y: -3},
			-3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Mul - Positive",
			args{x: 1, y: 1},
			1,
		},
		{
			"Mul - ByZero",
			args{x: 0, y: -3},
			0,
		},
		{
			"Mul - Negative",
			args{x: 1, y: -3},
			-3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Div - Positive",
			args{x: 1, y: 1},
			1,
		},
		{
			"Div - Negative",
			args{x: 1, y: -3},
			-3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Pow - Positive",
			args{x: 0, y: 1}, // 1^0
			1,
		},
		{
			"Pow - Negative",
			args{x: 1, y: 1}, // 1^1
			1,
		},
		{
			"Pow - Negative",
			args{x: 1, y: -3}, // -3^1
			-3,
		},
		{
			"Pow - BigNum",
			args{x: 8, y: 2}, // 2^8
			256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

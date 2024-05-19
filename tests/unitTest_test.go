package tests

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1+2=3", args: args{x: 1, y: 2}, want: 3},
		{name: "2+3=5", args: args{x: 2, y: 3}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevide(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "3/2=1.5", args: args{x: 3, y: 2}, want: 1.5},
		{name: "5/2=2.5", args: args{x: 5, y: 2}, want: 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Devide(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Devide() = %v, want %v", got, tt.want)
			}
		})
	}
}

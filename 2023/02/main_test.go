package main

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args args
		want Game
	}{
		{
			args: args{line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"},
			want: Game{
				ID: 1,
				Subsets: []Subset{
					{
						Red:   4,
						Green: 0,
						Blue:  3,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Red:   0,
						Green: 2,
						Blue:  0,
					},
				},
			},
		},
		{
			args: args{line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"},
			want: Game{
				ID: 3,
				Subsets: []Subset{
					{
						Red:   20,
						Green: 8,
						Blue:  6,
					},
					{
						Red:   4,
						Green: 13,
						Blue:  5,
					},
					{
						Red:   1,
						Green: 5,
						Blue:  0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.line, func(t *testing.T) {
			if got := NewGame(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

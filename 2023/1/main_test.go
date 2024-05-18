package main

import (
	"testing"
)

func Test_extractNumber(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Normal case",
			args: args{line: "1abc2"},
			want: 12,
		},
		{
			name: "Normal case 2",
			args: args{line: "pqr3stu8vwx"},
			want: 38,
		},
		{
			name: "Normal case 3",
			args: args{line: "a1b2c3d4e5f"},
			want: 15,
		},
		{
			name: "Single Digit",
			args: args{line: "treb7uchet"},
			want: 77,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractNumber(tt.args.line); got != tt.want {
				t.Errorf("extractNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_realExtractNumber(t *testing.T) {
	matchTree := NewStringMatchTree(NewMatchTreeArgs{
		{
			Word:   "one",
			Result: 1,
		},
		{
			Word:   "two",
			Result: 2,
		},
		{
			Word:   "three",
			Result: 3,
		},
		{
			Word:   "four",
			Result: 4,
		},
		{
			Word:   "five",
			Result: 5,
		},
		{
			Word:   "six",
			Result: 6,
		},
		{
			Word:   "seven",
			Result: 7,
		},
		{
			Word:   "eight",
			Result: 8,
		},
		{
			Word:   "nine",
			Result: 9,
		},
	})
	type args struct {
		line string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{line: "two1nine"},
			want: 29,
		},
		{
			args: args{line: "eightwothree"},
			want: 83,
		},
		{
			args: args{line: "xtwone3four"},
			want: 24,
		},
		{
			args: args{line: "zoneight234"},
			want: 14,
		},
		{
			args: args{line: "twonine"},
			want: 29,
		},
		{
			args: args{line: "fone2"},
			want: 12,
		},
		{
			args: args{line: "threight6"},
			want: 86,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.line, func(t *testing.T) {
			if got := realExtractNumber(tt.args.line, &matchTree); got != tt.want {
				t.Errorf("realExtractNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

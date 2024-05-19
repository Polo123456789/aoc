package main

import (
	"reflect"
	"testing"
)

func TestNewGraphCreationInstruction(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		args args
		want NodeCreationInstruction
	}{
		{
			args: args{line: "jqt: rhn xhk nvd"},
			want: NodeCreationInstruction{
				BaseID:       "jqt",
				ConnectToIDS: []string{"rhn", "xhk", "nvd"},
			},
		},
		{
			args: args{line: "xhk: hfx"},
			want: NodeCreationInstruction{
				BaseID:       "xhk",
				ConnectToIDS: []string{"hfx"},
			},
		},
		{
			args: args{line: "mxs: nqs qlm xfk pnf knr lmj bhp"},
			want: NodeCreationInstruction{
				BaseID: "mxs",
				ConnectToIDS: []string{
					"nqs",
					"qlm",
					"xfk",
					"pnf",
					"knr",
					"lmj",
					"bhp",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.line, func(t *testing.T) {
			if got := NewNodeCreationInstruction(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraphCreationInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

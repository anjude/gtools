package handler

import (
	"github.com/anjude/gtools/constant"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				args: []string{
					constant.HelpCmd.String(),
					"test",
				},
			},
		},
		{
			name: "",
			args: args{
				args: []string{
					constant.VersionCmd.String(),
					"test",
				},
			},
		},
		//{
		//	name: "",
		//	args: args{
		//		args: []string{
		//			constant.ChatCmd.String(),
		//			"你是谁",
		//		},
		//	},
		//},
		//{
		//	name: "",
		//	args: args{
		//		args: []string{
		//			constant.ShutdownCmd.String(),
		//			"test",
		//		},
		//	},
		//},
		{
			name: "",
			args: args{
				args: []string{
					constant.RandCmd.String(),
					"6",
					"number",
				},
			},
		},
		{
			name: "",
			args: args{
				args: []string{
					constant.RandCmd.String(),
					"20",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Execute(tt.args.args)
		})
	}
}

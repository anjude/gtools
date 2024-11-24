package rand

import "testing"

func TestHandler_Handle(t *testing.T) {
	type fields struct {
		Command string
		Desc    string
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				args: []string{"10", "string"},
			},
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				args: []string{"10", "number"},
			},
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				args: []string{"20", "all"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Handler{
				Command: tt.fields.Command,
				Desc:    tt.fields.Desc,
			}
			s.Handle(tt.args.args)
		})
	}
}

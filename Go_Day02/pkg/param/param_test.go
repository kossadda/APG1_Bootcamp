package param

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		args []string
		exp  Param
		err  error
	}{
		{
			name: "use -sl",
			args: []string{"-sl", "/home"},
			exp:  Param{Path: "/home", flags: slMask},
			err:  nil,
		},
		{
			name: "use -d",
			args: []string{"-d", "/home"},
			exp:  Param{Path: "/home", flags: dMask},
			err:  nil,
		},
		{
			name: "use -f",
			args: []string{"-f", "/home"},
			exp:  Param{Path: "/home", flags: fMask},
			err:  nil,
		},
		{
			name: "use -f with -ext",
			args: []string{"-f", "-ext", "txt", "/home"},
			exp:  Param{Path: "/home", Ext: "txt", flags: fMask | extMask},
			err:  nil,
		},
		{
			name: "use -ext without -f",
			args: []string{"-ext", "txt", "/home"},
			exp:  Param{},
			err:  errors.New("flag -ext provided but -f is not used"),
		},
		{
			name: "use -ext with empty value",
			args: []string{"-f", "-ext", "", "/home"},
			exp:  Param{},
			err:  errors.New("flag -ext provided but extension is empty"),
		},
		{
			name: "use multiple flags",
			args: []string{"-sl", "-f", "-ext", "log", "/home"},
			exp:  Param{Path: "/home", Ext: "log", flags: slMask | fMask | extMask},
			err:  nil,
		},
		{
			name: "no flags with path",
			args: []string{"/home"},
			exp:  Param{Path: "/home"},
			err:  nil,
		},
		{
			name: "no path provided",
			args: []string{"-sl"},
			exp:  Param{},
			err:  errors.New("provide one path argument at the end"),
		},
		{
			name: "too many arguments",
			args: []string{"-sl", "/home", "/extra"},
			exp:  Param{},
			err:  errors.New("provide one path argument at the end"),
		},
		{
			name: "unknown flag",
			args: []string{"-unknown", "/home"},
			exp:  Param{},
			err:  errors.New("flag provided but not defined: -unknown"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.name, tt.args)
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("New() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.exp {
				t.Errorf("New() = %v, want %v", got, tt.exp)
			}
		})
	}
}

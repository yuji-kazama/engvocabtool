/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewAddCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
		wantErr bool
	}{
		{
			name: "normal",
			args: []string{"add", "cool"},
			want: "Success: word has been added",
			wantErr: false,
		},
		// {
		// 	name: "no arg",
		// 	args: []string{"add"},
		// 	want: "accepts 1 arg(s), received 0",
		// 	wantErr: true,
		// },
		// {
		// 	name: "already added",
		// 	args: []string{"add", "test"},
		// 	want: "already exists",
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := NewRootCmd()
			cmd.SetOut(buf)
			cmd.SetArgs(tt.args)

			if err := cmd.Execute(); err != nil {
				if tt.wantErr {
					if tt.want != err.Error() {
						t.Errorf("unexpected error response: error = %v, got = %v", tt.want, err.Error())
					}
					return
				}
			}

			got := buf.String()
			if !strings.HasPrefix(got, tt.want) {
				t.Errorf("unexpected response: want = %v, got = %v", tt.want, got)
			}
			
		})
	}
}

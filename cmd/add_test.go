/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("NOTION_DATABASE_ID", "b3a02fe5c7b14208a880cdd92e5b10ae")
	status := m.Run()
	os.Exit(status)
}
func TestNewAddCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name: "normal",
			args: []string{"add", "cool"},
			want: "Success: word has been added",
			wantErr: false,
		},
		{
			name: "no arg",
			args: []string{"add"},
			want: "accepts 1 arg(s), received 0",
			wantErr: true,
		},
		{
			name: "already added",
			args: []string{"add", "test"},
			want: "already exists",
			wantErr: true,
		},
		{
			name:    "no definition",
			args:    []string{"add", "hogaegae"},
			want:    "no matching word was found",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := NewRootCmd()
			cmd.SetOut(buf)
			cmd.SetArgs(tt.args)

			err := cmd.Execute()
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error = %#v", err)
			}
			got := buf.String()
			if !strings.HasPrefix(got, tt.want) {
				t.Fatalf("want %#v, but %#v", tt.want, got)
			}
		})
	}
}

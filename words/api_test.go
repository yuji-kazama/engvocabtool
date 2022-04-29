package words

import (
	"testing"
)

func TestClient_GetEverything(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args {
				word: "friend",
			},
			wantErr: false, 
		},
		{
			name: "unknown",
			args: args {
				word: "awhoefiuawef",
			},
			wantErr: true, 
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.GetEverything(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetEverything() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				t.Logf("Word: %v", got.Word)
				t.Logf("Frequency: %v", got.Frequency)
			}
		})
	}
}

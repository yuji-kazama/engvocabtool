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
		want    AllResults
		wantErr bool
		err error
	}{
		{
			name: "normal",
			args: args {
				word: "friend",
			},
			want: AllResults{
				Word: "friend",
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

			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
			if !tt.wantErr && got.Word != tt.want.Word {
				t.Fatalf("want %#v, but %#v", got.Word, tt.want.Word)
			}
		})
	}
}

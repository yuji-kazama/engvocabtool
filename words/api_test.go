package words

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) * http.Client {
	return &http.Client {
		Transport: fn,
	}
}

func newMockedClient(t *testing.T, requestMockFile string, statusCode int) *http.Client {
	return newTestClient(func(req * http.Request) * http.Response {
		b, err := os.Open(requestMockFile)
		if err != nil {
			t.Fatal(err)
		}
		resp := &http.Response{
			StatusCode: statusCode,
			Body: b,
			Header: make(http.Header),
		}
		return resp
	})
}

func TestClient_GetEverythingWithMock(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
		err error
		filePath string
		statusCode int
	}{
		{
			name: "normal",
			args: args {
				word: "apple",
			},
			want: &Response{
				Word: "apple",
				Results: []Result {
					{
						Definition: "native Eurasian tree widely cultivated in many varieties for its firm rounded edible fruits",
						PartOfSpeech: "noun",
						Synonyms: []string{
							"malus pumila",
							"orchard apple tree",
						},
					},
					{
						Definition: "fruit with red or yellow or green skin and sweet to tart crisp whitish flesh",
						PartOfSpeech: "noun",
					},
				},
				Frequency: 4.34,
				Pronunciation: Pronunciation {
					All: "'æpəl",
				},
			},
			wantErr: false, 
			filePath: "testdata/getEverything.json",
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newMockedClient(t, tt.filePath, tt.statusCode)
			client := NewClient(WithHttpClient(c))

			got, err := client.GetEverything(context.Background(), tt.args.word)

			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
			if d := cmp.Diff(got, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}

}

func TestClient_GetEverything(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
		err error
	}{
		{
			name: "normal",
			args: args {
				word: "friend",
			},
			want: Response{
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
			err: fmt.Errorf("error code 404: your request is invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient()
			got, err := client.GetEverything(context.Background(), tt.args.word)

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

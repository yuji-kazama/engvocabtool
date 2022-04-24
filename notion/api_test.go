package notion

import (
	"testing"
)

func TestClient_GetPage(t *testing.T) {
	type args struct {
		pageId string
	}
	type page struct {
		Object string
		Id     string
	}
	tests := []struct {
		name    string
		args    args
		want    page
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				pageId: "25ec2f4cdb444dc6aee5ebab57d5cd91",
			},
			want: page{
				Object: "page",
				Id:     "25ec2f4cdb444dc6aee5ebab57d5cd91",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.GetPage(tt.args.pageId)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Client.GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Object != tt.want.Object {
				t.Fatalf("Client.GetPage() = %v, want %v", got, tt.want)
			}
			// t.Logf("\n Name = %v", got.Properties.Name.Title[0].Text.Content)
			// t.Logf("\n Frequency = %v", got.Properties.Frequency.Number)
		})
	}
}

func TestClient_UpdatePage(t *testing.T) {
	type args struct {
		pageId string
		item   *Item
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				pageId: "36ad047fbf4d41879eb90cc028ea7074",
				item: &Item{
					Properties: struct {
								Frequency struct {
									Number int "json:\"number\""
								} "json:\"Frequency\""
							}{
								Frequency: struct {
									Number int "json:\"number\""
								}{
									Number: 0,
								},
							},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			if err := c.UpdatePage(tt.args.pageId, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdatePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

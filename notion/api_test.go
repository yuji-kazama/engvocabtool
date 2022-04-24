package notion

import (
	"os"
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
				pageId: "36ad047fbf4d41879eb90cc028ea7074",
			},
			want: page{
				Object: "page",
				Id:     "36ad047fbf4d41879eb90cc028ea7074",
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
		json   string
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
				json: `{
					"properties": {
						"Frequency": {
							"number": 9
						},
						"Meaning": {
							"rich_text": [
								{
									"text": {
										"content": "hoge"
									}
								}
							]
						}
					} 
				}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			if err := c.UpdatePage(tt.args.pageId, tt.args.json); (err != nil) != tt.wantErr {
				t.Errorf("Client.UpdatePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_PostPage(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				json: `{
					"parent": {
						"database_id": "` + os.Getenv("NOTION_DATABASE_ID") + `"
					},
					"properties": {
						"Name": {
							"title": [
								{
									"text": {
										"content": "This is test from API"
									}
								}
							]
						}
					} 
				}`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			if err := c.PostPage(tt.args.json); (err != nil) != tt.wantErr {
				t.Errorf("Client.PostPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_GetAllPages(t *testing.T) {
	type result struct {
		Object string
	}
	tests := []struct {
		name    string
		want    result
		wantErr bool
	}{
		{
			name: "normal",
			want: result{
				Object: "list",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.GetAllPages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAllPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Object != tt.want.Object {
				t.Errorf("Client.GetAllPages() = %v, want %v", got, tt.want)
		}
		})
	}
}

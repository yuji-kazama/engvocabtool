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
						"Status": {
							"select": {
								"name": "2: Recognized"
							}
						},
						"Check Num": {
							"select": {
								"name": "1"
							}
						},
						"Study Date": {
							"date": {
								"start": "2022-04-09",
								"end": null,
								"time_zone": null
							}
						},
						"Class": {
							"select": {
								"name": "Adjective"
							}
						},
						"Frequency": {
							"number": 4
						},
						"Meaning": {
							"rich_text": [
								{
									"text": {
										"content": "updated"
									}
								}
							]
						},
						"Sentence": {
							"rich_text": [
								{
									"text": {
										"content": "updated"
									}
								}
							]
						},
						"Note": {
							"rich_text": [
								{
									"text": {
										"content": "updated"
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
			got, err := c.PostPage(tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.PostPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got.URL)
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
			name:    "normal",
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

func TestClient_GetPageByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				name: "test",
			},
			wantErr: false,
		},
		{
			name: "not exist",
			args: args{
				name: "afjawifaiwe",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.GetPageByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetPageByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Result: %v", got.Results)
		})
	}
}

func TestClient_Exist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		args   args
		want   bool
	}{
		{
			name: "not exist",
			args: args{
				name: "afjawifaiwe",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			if got := c.Exist(tt.args.name); got != tt.want {
				t.Errorf("Client.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

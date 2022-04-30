package notion

import (
	"errors"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("NOTION_DATABASE_ID", "b3a02fe5c7b14208a880cdd92e5b10ae")
	status := m.Run()
	os.Exit(status)
}

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
		err error
	}{
		{
			name: "normal",
			args: args{
				pageId: "781b31df08284c6ab414e4f487fa60e0",
			},
			want: page{
				Object: "page",
				Id:     "781b31df08284c6ab414e4f487fa60e0",
			},
			wantErr: false,
		},
		{
			name: "empty args",
			args: args{
				pageId: "",
			},
			wantErr: true,
			err: errors.New("response error: 400"),
		},
		{
			name: "no existing pageId",
			args: args{
				pageId: "no-exting-page",
			},
			wantErr: true,
			err: errors.New("response error: 400"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.GetPage(tt.args.pageId)
			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
			if !tt.wantErr && got.Object != tt.want.Object {
				t.Fatalf("want %#v, but %#v", got.Object, tt.want.Object)
			}
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
		err error
	}{
		{
			name: "normal: Status",
			args: args{
				pageId: "781b31df08284c6ab414e4f487fa60e0",
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
								"start": "2022-04-19",
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
							"number": 1
						},
						"Meaning": {
							"rich_text": [
								{
									"text": {
										"content": "updated2"
									}
								}
							]
						},
						"Sentence": {
							"rich_text": [
								{
									"text": {
										"content": "updated2"
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
			_, err := c.UpdatePage(tt.args.pageId, tt.args.json)

			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
		})
	}
}

func TestClient_CreatePage(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err error
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
										"content": "TEST_FROM_API"
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
			_, err := c.CreatePage(tt.args.json)

			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
		})
	}
}

func TestClient_GetAllPages(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		err  error
	}{
		{
			name:    "normal",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			_, err := c.GetAllPages()
			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
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
		err error
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
			_, err := c.GetPageByName(tt.args.name)
			if !tt.wantErr && err != nil{
				t.Fatalf("unexpected error = %#v", err)
			}
			if tt.wantErr && err == tt.err {
				t.Fatalf("want %#v, but %#v", tt.err, err)
			}
		})
	}
}

func TestClient_Exist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
		wantErr bool
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
			got := c.Exist(tt.args.name)

			if !tt.wantErr && got != tt.want {
				t.Fatalf("want %#v, but %#v", got, tt.want)
			}
		})
	}
}

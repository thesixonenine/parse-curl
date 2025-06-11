package parse_curl

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		curl string
	}
	tests := []struct {
		name    string
		args    args
		want    *Request
		wantErr bool
	}{
		{
			name: "GET simple",
			args: args{
				curl: "curl 'https://example.com'",
			},
			want: &Request{
				Method: "GET",
				Url:    "https://example.com",
				Header: map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "GET with method",
			args: args{
				curl: "curl --location --request GET 'https://example.com'",
			},
			want: &Request{
				Method: "GET",
				Url:    "https://example.com",
				Header: map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "GET with header and cookie",
			args: args{
				curl: "curl 'https://example.com' \\\n --cookie 'a=b;c=d' -H 'dnt: 1' -H 'Accept: text/plain' --header 'User-Agent: Chrome'",
			},
			want: &Request{
				Method: "GET",
				Url:    "https://example.com",
				Header: map[string]string{"Cookie": "a=b;c=d", "dnt": "1", "Accept": "text/plain", "User-Agent": "Chrome"},
			},
			wantErr: false,
		},
		{
			name: "POST",
			args: args{
				curl: "curl -d 'foo=bar' https://example.com",
			},
			want: &Request{
				Method: "POST",
				Url:    "https://example.com",
				Header: map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
				Body:   "foo=bar",
			},
			wantErr: false,
		},
		{
			name: "PUT",
			args: args{
				curl: "curl -X PUT 'https://example.com'",
			},
			want: &Request{
				Method: "PUT",
				Url:    "https://example.com",
				Header: map[string]string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.curl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

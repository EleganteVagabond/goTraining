package sitemap

import (
	"encoding/xml"
	"testing"
)

func TestEncodeListToXML(t *testing.T) {
	type args struct {
		xmlEnc *xml.Encoder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EncodeListToXML(tt.args.xmlEnc)
		})
	}
}

func Test_getFullPathURL(t *testing.T) {
	type args struct {
		urlIn    string
		basePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple test",
			args: args{urlIn: "http://go.com", basePath: "http://go.com"},
			want: "http://go.com",
		},
		{
			name: "absolute path",
			args: args{urlIn: "/", basePath: "http://go.com"},
			want: "http://go.com/",
		},
		{
			name: "bookmark",
			args: args{urlIn: "#food", basePath: "http://go.com"},
			want: "http://go.com/#food",
		},
		{
			name: "relative path",
			args: args{urlIn: "goodies/bob", basePath: "http://go.com"},
			want: "http://go.com/goodies/bob",
		},
		{
			name: "actual html page title",
			args: args{urlIn: "index.html", basePath: "http://go.com"},
			want: "http://go.com/index.html",
		},
		{
			name: "email",
			args: args{urlIn: "@mailto:someone@someone.org", basePath: "http://go.com"},
			want: "http://go.com/@mailto:someone@someone.org",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFullPathURL(tt.args.urlIn, tt.args.basePath); got != tt.want {
				t.Errorf("getFullPathURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInDomain(t *testing.T) {
	type args struct {
		url    string
		domain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple",
			args: args{
				url:    "http://www.goo.com",
				domain: "www.goo.com",
			},
			want: true,
		},
		{
			name: "long path",
			args: args{
				url:    "http://www.goo.com/food/for/thought/yum.xhtml",
				domain: "www.goop.com",
			},
			want: false,
		},
		{
			name: "false",
			args: args{
				url:    "http://www.goo.com",
				domain: "www.goop.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInDomain(tt.args.url, tt.args.domain); got != tt.want {
				t.Errorf("isInDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopulateSitemap(t *testing.T) {
	tests := []struct {
		name     string
		startLoc string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PopulateSitemap(tt.startLoc)
		})
	}
}

func Test_extractSiteInfo(t *testing.T) {
	type want struct {
		startLoc string
		siteroot string
		domain   string
	}
	tests := []struct {
		name     string
		startLoc string
		want     want
	}{
		{
			name:     "simple",
			startLoc: "http://test.com",
			want: want{
				startLoc: "http://test.com",
				siteroot: "http://test.com",
				domain:   "test.com",
			},
		},
		{
			name:     "no transport",
			startLoc: "test.com",
			want: want{
				startLoc: "http://test.com",
				siteroot: "http://test.com",
				domain:   "test.com",
			},
		},
		{
			name:     "somewhere in the middle of the site, no transport",
			startLoc: "test.com/booger/head/poop",
			want: want{
				startLoc: "http://test.com/booger/head/poop",
				siteroot: "http://test.com",
				domain:   "test.com",
			},
		},
		{
			name:     "somewhere in the middle of the site, with transport",
			startLoc: "http://test.com/booger/head/poop",
			want: want{
				startLoc: "http://test.com/booger/head/poop",
				siteroot: "http://test.com",
				domain:   "test.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if s1, s2, s3 := extractSiteInfo(tt.startLoc); s1 != tt.want.startLoc || s2 != tt.want.siteroot || s3 != tt.want.domain {
				t.Errorf("isInDomain() = %v, want %v", want{s1, s2, s3}, tt.want)
			}
		})
	}
}

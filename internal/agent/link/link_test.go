package link

import (
	"ShortLand/internal/model"
	"reflect"
	"testing"
	"time"
)

func TestCreateShortLink(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want model.LinkTable
	}{
		{
			name: "success",
			args: args{"https://github.com/gorilla/mux"},
			want: model.LinkTable{
				OriginLink: "https://github.com/gorilla/mux",
				ShortLink:  "Lpny6Wyanx",
				ExpireAt:   time.Now().Add(5 * time.Minute).Unix(),
			},
		},
		{
			name: "error",
			args: args{"https://github.com/gorilla/mux"},
			want: model.LinkTable{
				OriginLink: "https://github.com/gorilla/mux",
				ShortLink:  "Lpny5Wyanx",
				ExpireAt:   time.Now().Add(5 * time.Minute).Unix(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateShortLink(tt.args.link); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateShortLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

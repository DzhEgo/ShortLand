package link

import (
	"ShortLand/internal/portal/service/link/typeStorage"
	"testing"
)

func Test_linkService_createShortLink(t *testing.T) {
	type fields struct {
		stor StorageLink
	}
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "create short link",
			fields: fields{
				stor: typeStorage.NewInMemory(),
			},
			args: args{
				link: "http://google.com",
			},
			want:    "4K9rajbFjd",
			wantErr: false,
		},
		{
			name: "create empty short link",
			fields: fields{
				stor: typeStorage.NewInMemory(),
			},
			args: args{
				link: "",
			},
			want:    "4K9rajbFjd",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &linkService{
				stor: tt.fields.stor,
			}
			got, err := s.CreateShortLink(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("createShortLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createShortLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

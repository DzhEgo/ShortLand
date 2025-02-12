package app

import "testing"

func TestStartApp(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StartApp(); (err != nil) != tt.wantErr {
				t.Errorf("StartApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

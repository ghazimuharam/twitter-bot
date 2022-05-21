package resolver

import "testing"

func TestGetRedirectURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "failed resolve url",
			args: args{
				url: "https://--google.com",
			},
			wantErr: true,
		},
		{
			name: "success resolve url",
			args: args{
				url: "https://google.com",
			},
			want: "https://www.google.com/",
		},
		{
			name: "success resolve url twitter",
			args: args{
				url: "https://t.co/RF0GcjnZF0",
			},
			want: "https://twitter.com/numberoneeisimp/status/1527397191165153282/photo/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRedirectURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRedirectURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRedirectURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

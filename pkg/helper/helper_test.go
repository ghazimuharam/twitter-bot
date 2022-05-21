package helper

import "testing"

func TestTwitterURLBuilder(t *testing.T) {
	type args struct {
		handler string
		tweetID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success return url",
			args: args{
				handler: "GAMEFESS_",
				tweetID: "1527397191165153282",
			},
			want: "https://twitter.com/GAMEFESS_/status/1527397191165153282",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwitterURLBuilder(tt.args.handler, tt.args.tweetID); got != tt.want {
				t.Errorf("TwitterURLBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

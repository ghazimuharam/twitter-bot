package regex

import "testing"

func TestRemoveURLFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "without link",
			args: args{
				str: "this is a test regex val!",
			},
			want: "this is a test regex val!",
		},
		{
			name: "with link",
			args: args{
				str: " this is a test regex val! https://t.co/1B5ci0KIN6",
			},
			want: "this is a test regex val!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveURLFromString(tt.args.str); got != tt.want {
				t.Errorf("RemoveURLFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

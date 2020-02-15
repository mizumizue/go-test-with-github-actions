package main

import "testing"

func TestGreeting(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{name: "trewanek"},
			want: "Hello, trewanek",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greeting(tt.args.name); got != tt.want {
				t.Errorf("Greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

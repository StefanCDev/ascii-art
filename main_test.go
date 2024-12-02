package main
import (
	"reflect"
	"testing"
)
func TestRowFinder(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RowFinder(tt.args.b); got != tt.want {
				t.Errorf("RowFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAsciiArt(t *testing.T) {
	type args struct {
		input string
		font  []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AsciiArt(tt.args.input, tt.args.font)
		})
	}
}
func TestGetfont(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Getfont(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Getfont() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
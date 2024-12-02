package main
import (
	"fmt"
	"os"
	"strings"
)
// finds the row where the character we want to print start from in the text file
func RowFinder(b byte) int {
	return (int(b)-32)*9 + 1
}
func AsciiArt(input string, font []string) {
	if input == "" {
		return
	}
	if input == `\n` {
		fmt.Println()
		return
	}
	// split the input by new line whenever it finds \n
	lines := strings.Split(input, `\n`)
	// looping through the slice that contains the arguments
	for i := 0; i < len(lines); i++ {
		switch lines[i] {
		case "":
			fmt.Println()
		default:
			for j := 0; j < 8; j++ {
				for k := 0; k < len(lines[i]); k++ {
					//fmt.Println(j,k)
					fmt.Print(font[RowFinder(lines[i][k])+j])
				}
				fmt.Println()
			}
		}
	}
}
func Getfont() []string {
	fontFile, _ := os.ReadFile(os.Args[2])
	// if len(os.Args) == 3 {
	// 	fontFile, err = os.ReadFile(os.Args[2] + ".txt")
	return strings.Split(string(fontFile), "\n")
}
func main() {
	font := Getfont()
	input := os.Args[1]
	AsciiArt(input, font)
}
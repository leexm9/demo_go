package main

import "fmt"

func main() {
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
	fmt.Println()

	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}
}

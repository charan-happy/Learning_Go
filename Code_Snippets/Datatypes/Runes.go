/*
> Rune type is the same as int32 which constitutes unicode points
> Runes can be created using enclosing one or more characters in single quotes like 'g', '\t' etc. In between single quotes except for a new line and an unescaped single quote.
*/

package main

import "fmt"

func main() {
	arr1 := 'A'
	arr := '✅'
	fmt.Println("Character %c, Unicode %U", arr1, arr1)
	fmt.Println("\nCharacter %c, Unicode %U", arr, arr)
}

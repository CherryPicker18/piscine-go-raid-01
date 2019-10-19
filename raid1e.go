package student

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	if y == 1 {
		printTopE(x)
	} else {
		printTopE(x)
		for i := 0; i < y-2; i++ {
			printMidE(x)
		}
		printBotE(x)
	}
}

func printTopE(x int) {
	if x == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('A')
		for j := 0; j < x-2; j++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
	}
}
func printMidE(x int) {
	if x == 1 {
		z01.PrintRune('B')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('B')
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('B')
		z01.PrintRune('\n')
	}
}
func printBotE(x int) {
	if x == 1 {
		z01.PrintRune('C')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('C')
		for j := 0; j < x-2; j++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
		z01.PrintRune('\n')
	}
}

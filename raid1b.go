package student

import "github.com/01-edu/z01"

func Raid1b(x, y int) {
	if y == 1 {
		printTopB(x)
	} else {
		printTopB(x)
		for i := 0; i < y-2; i++ {
			printMidB(x)
		}
		printBotB(x)
	}
}

func printTopB(x int) {
	if x == 1 {
		z01.PrintRune('/')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('/')
		for j := 0; j < x-2; j++ {
			z01.PrintRune('*')
		}
		z01.PrintRune(92)
		z01.PrintRune('\n')
	}
}
func printMidB(x int) {
	if x == 1 {
		z01.PrintRune('*')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('*')
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('*')
		z01.PrintRune('\n')
	}
}
func printBotB(x int) {
	if x == 1 {
		z01.PrintRune(92)
		z01.PrintRune('\n')
	} else {

		z01.PrintRune(92)
		for j := 0; j < x-2; j++ {
			z01.PrintRune('*')
		}
		z01.PrintRune('/')
		z01.PrintRune('\n')
	}
}

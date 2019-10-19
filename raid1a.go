package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if y == 1 {
		printEdge(x)
	} else {
		printEdge(x)
		for i := 0; i < y-2; i++ {
			printMid(x)
		}
		printEdge(x)
	}
}

func printEdge(x int) {
	if x == 1 {
		z01.PrintRune('o')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('o')
		for j := 0; j < x-2; j++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
	}
}
func printMid(x int) {
	if x == 1 {
		z01.PrintRune('|')
		z01.PrintRune('\n')
	} else {

		z01.PrintRune('|')
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('|')
		z01.PrintRune('\n')
	}
}

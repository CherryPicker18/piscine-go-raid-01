package student

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	if y == 1 {
		printEdgeD(x)
	} else {
		printEdgeD(x)
		for i := 0; i < y-2; i++ {
			printMidD(x)
		}
		printEdgeD(x)
	}
}

func printEdgeD(x int) {
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
func printMidD(x int) {
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

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	first := atoi(args[0])
	second := atoi(args[1])

	QuadE(first, second)

}

func atoi(s string) int {
	sd := []rune(s)
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(rune(sd[i] - 48))

	}
	return res
}

func QuadA(x, y int) {
	c1 := 111 // o
	c2 := 111 // o
	yD := 124 // |
	xD := 45  // -
	abc := 1
	XxY(x, y, c1, c2, yD, xD, abc)
}

func QuadB(x, y int) {
	c1 := 47 // /
	c2 := 92 // \
	yD := 42 // *
	xD := 42 // *
	abc := 2
	XxY(x, y, c1, c2, yD, xD, abc)
}

func QuadC(x, y int) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 3
	XxY(x, y, c1, c2, yD, xD, abc)
}

func QuadD(x, y int) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 1
	XxY(x, y, c1, c2, yD, xD, abc)
}

func QuadE(x, y int) {
	c1 := 65 // A
	c2 := 67 // C
	yD := 66 // B
	xD := 66 // B
	abc := 2
	XxY(x, y, c1, c2, yD, xD, abc)
}

func XxY(x, y, c1, c2, yD, xD, abc int) {
	if x < 1 || y < 1 {
		return
	}
	delX := x - 2
	delY := y - 2
	switch abc {
	case 1:
		xaxys(x, delX, c1, c2, xD)
		for i := 0; i < delY; i++ {
			yaxys(y, delX, x, yD)
		}
		if y > 1 {
			xaxys(x, delX, c1, c2, xD)
		}
	case 2:
		xaxys(x, delX, c1, c2, xD)
		for i := 0; i < delY; i++ {
			yaxys(y, delX, x, yD)
		}
		if y > 1 {
			xaxys(x, delX, c2, c1, xD)
		}
	case 3:
		xaxys(x, delX, c1, c1, xD)
		for i := 0; i < delY; i++ {
			yaxys(y, delX, x, yD)
		}
		if y > 1 {
			xaxys(x, delX, c2, c2, xD)
		}

	}
}

func xaxys(x, delX, c1, c2, xD int) {
	z01.PrintRune(rune(c1))
	for delixx := 0; delixx < delX; delixx++ {
		z01.PrintRune(rune(xD))
	}
	if x > 1 {
		z01.PrintRune(rune(c2))
	}
	z01.PrintRune('\n')
}

func yaxys(y, delX, x, yD int) {
	z01.PrintRune(rune(yD))
	for delixx := 0; delixx < delX; delixx++ {
		z01.PrintRune(' ')
	}
	if x > 1 {
		z01.PrintRune(rune(yD))
	}
	z01.PrintRune('\n')
}

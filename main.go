package main

import "fmt"

func main() {
	s := "DDUUDDUDUUUD"
	sl := 0
	valley := 0
	prevsl := 0

	for i := 0; i < len(s); i++ {
		prevsl = sl

		if s[i] == 'D' {
			sl--
		} else {
			sl++
		}
		if (sl == -1) && (prevsl == 0) {
			valley++
		}
	}

	fmt.Println(valley)
}

package main

import (
	"fmt"
	"strconv"
)

var chess [8][8]int
var ans []string
var ansforchess [][8]int

func main() {
	var ind int
	fmt.Scan(&ind)
	var qP [8]int
	queenPutter(qP, 0)
	fmt.Println(ind)
	for i, k := range ans {
		if i == ind {
			fmt.Println(k)
		}
	}
	printQueen(ind)
}
func transform(a [8]int) string {
	alf := "abcdefjh"
	t := ""
	for i, k := range a {
		t = t + string(alf[i]) + strconv.Itoa(k) + " "
	}
	return t
}
func queenPutter(qp [8]int, i int) {
	if i == 8 {
		ans = append(ans, transform(qp))
		ansforchess = append(ansforchess, qp)
		return
	}
	for j := 0; j < 8; j++ {
		if isItSafe(i, j) {
			qp[i] = j
			chess[i][j] = 1
			queenPutter(qp, i+1)
		}
		chess[i][j] = 0
	}
}
func isItSafe(i, j int) bool {
	if col(i, j) == false || rDiag(i, j) == false || lDiag(i, j) == false {
		return false
	}
	return true
}
func col(i, j int) bool {
	if i == 0 {
		if chess[i][j] != 0 {
			return false
		} else {
			return true
		}
	}
	if chess[i][j] != 0 {
		return false
	}
	return col(i-1, j)
}
func rDiag(i, j int) bool {
	if i == 0 || j == 7 {
		if chess[i][j] != 0 {
			return false
		} else {
			return true
		}
	}
	if chess[i][j] != 0 {
		return false
	}
	return rDiag(i-1, j+1)
}
func lDiag(i, j int) bool {
	if i == 0 || j == 0 {
		if chess[i][j] != 0 {
			return false
		} else {
			return true
		}
	}
	if chess[i][j] != 0 {
		return false
	}
	return lDiag(i-1, j-1)
}
func printQueen(a int) {
	var n1, n2 int
	st := "|"
	for k, l := range ansforchess[a] {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if n1 == 8 {
					n1 = 0
					n2 += 1
					if checkforNewLine(st) {
						st = st + "\n" + "|"
					}

				}
				n1 += 1
				st = collect(k, l, i, j, n1, n2, st)
			}
		}

	}
	fmt.Println(st)
}
func checkforNewLine(st string) bool {
	if st[len(st)-2] == '\n' && st[len(st)-1] == '|' {
		return false
	}
	return true
}
func collect(k, l, i, j, n1, n2 int, st string) string {
	if k == i && l == j {
		st = st + "Q"
		if st[len(st)-1] != '|' {
			st = st + "|"
		}
	} else if k == i && l != j {
		if n2%2 == 0 {
			if n1%2 == 0 {
				st = st + " "
			} else {
				st = st + "#"
			}
		} else {
			if n1%2 == 0 {
				st = st + "#"
			} else {
				st = st + " "
			}
		}
		if st[len(st)-1] != '|' {
			st = st + "|"
		}
	}

	return st
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
2
3 3
000
000
000
111
111
111
4 4
0000
0000
0000
0000
2100
1200
0012
0021
4 4
1020
1200
1210
0000
0000
1200
2200
0000
3 3
012
012
012
010
111
011
8 8
00000000
00000000
00000000
00000000
00000000
00000000
00000000
10000000
00000000
01200000
02010000
00102000
00020100
00001020
00000210
10000000
2 7
0000000
0000000
2220111
0111222
2 7
0000000
0100010
2220111
1210202
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	var t int
	fmt.Sscanf(scanner.Text(), "%d", &t)

	for ; t > 0; t-- {
		scanner.Scan()
		var n, m int
		fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

		a := make([][]int, n)
		b := make([][]int, n)

		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
			scanner.Scan()
			line := scanner.Text()
			for j := 0; j < m; j++ {
				a[i][j] = int(line[j] - '0')
			}
		}

		for i := 0; i < n; i++ {
			b[i] = make([]int, m)
			scanner.Scan()
			line := scanner.Text()
			for j := 0; j < m; j++ {
				b[i][j] = int(line[j] - '0')
			}
		}

		if canTrans(a, b, n, m) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func canTrans(a, b [][]int, n, m int) bool {
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			if a[i][j] != b[i][j] {
				diff := (b[i][j] - a[i][j] + 3) % 3
				a[i][j] = (a[i][j] + diff) % 3
				a[i+1][j] = (a[i+1][j] + (2 * diff)) % 3
				a[i][j+1] = (a[i][j+1] + (2 * diff)) % 3
				a[i+1][j+1] = (a[i+1][j+1] + diff) % 3
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

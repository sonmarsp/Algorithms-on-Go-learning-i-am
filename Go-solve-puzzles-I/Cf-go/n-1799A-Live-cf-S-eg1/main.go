package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	testCase := readInt(in)
	fmt.Println("testCase")
	fmt.Println(testCase)

	for t := 0; t < testCase; t++ {
		input := readArrInt(in)
		fmt.Println("input")
		fmt.Println(input)

		query := readArrInt(in)
		fmt.Println("query")
		fmt.Println(query)

		result := solve(input[0], input[1], query)

		sb := strings.Builder{}

		for i := range result {
			if i != 0 {
				sb.WriteByte(' ')
			}

			sb.WriteString(strconv.Itoa(result[i]))
		}

		fmt.Println(sb.String())
	}

}

func solve(n, m int, query []int) []int {
	result := make([]int, n+m-1)
	//fmt.Println("result")
	//fmt.Println(result)

	linkedList := list.New()
	index := map[int]*list.Element{}

	for i := 1; i <= n; i++ {
		index[i] = linkedList.PushBack(i)
		result[i] = -1
		fmt.Println("result")
		fmt.Println(result)
	}

	for i, num := range query {
		if _, ok := index[num]; !ok {
			deleted := linkedList.Remove(linkedList.Back()).(int)
			fmt.Println("deleted")
			fmt.Println(deleted)

			if result[deleted] == -1 {
				result[deleted] = i + 1
			}

			delete(index, deleted)
			index[num] = linkedList.PushFront(num)

		} else {

			linkedList.MoveToFront(index[num])
		}
	}

	return result[1 : n+1]
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readString(in *bufio.Reader) string {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	return nStr
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

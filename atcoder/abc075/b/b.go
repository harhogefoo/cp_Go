package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var sc = bufio.NewScanner(os.Stdin)

// 文字列を1行取得
func strStdin() (stringInput string) {
	sc.Scan()
	stringInput = sc.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

// 整数値1つ取得
func intStdin() (int, error) {
	stringInput := strStdin()
	return strconv.Atoi(strings.TrimSpace(stringInput))
}

// 空白や空文字だけの値を除去したSplit()
func splitWithoutEmpty(stringTargeted string, delim string) (stringReturned []string) {
	stringSplited := strings.Split(stringTargeted, delim)

	for _, str := range stringSplited {
		if str != "" {
			stringReturned = append(stringReturned, str)
		}
	}
	return
}

// デリミタで分割して文字列スライスを取得
func splitStrStdin(delim string) (stringReturned []string) {
	// 文字列で1行取得
	stringInput := strStdin()

	// 空白で分割
	stringReturned = splitWithoutEmpty(stringInput, delim)

	return
}

// デリミタで分割して整数値スライスを取得
func splitIntStdin(delimiter string) (intAry []int) {
	// 文字列スライスを取得
	strAry := splitStrStdin(delimiter)

	// 整数スライスに保存
	for i := range strAry {
		num, err := strconv.Atoi(strAry[i])
		if err != nil {
			panic(err)
		}
		intAry = append(intAry, num)
	}
	return
}


func main() {
	action := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
	num := splitIntStdin(" ")
	h := num[0]

	matrix := make([][]string, h)
	for i := 0; i < h; i++ {
		matrix[i] = splitStrStdin("")
	}

	for i := 0; i < h; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "#" {
				continue
			}
			var count int
			for _, a := range action {
				x, y := i + a[0], j + a[1]
				if x < 0 || len(matrix) <= x {
					continue
				}
				if y < 0 || len(matrix[i]) <= y {
					continue
				}
				if matrix[x][y] == "#" {
					count++
				}
			}
			matrix[i][j] = strconv.Itoa(count)
		}
	}

	for _, m := range matrix {
		fmt.Println(strings.Join(m, ""))
	}
}

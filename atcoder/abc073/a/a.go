package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	n := splitStrStdin("")
	if n[0] == "9" || n[1] == "9" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

/* template */
var sc = bufio.NewScanner(os.Stdin)

// 文字列を1行取得
func strStdin() (stringInput string) {
	sc.Scan()
	stringInput = sc.Text()
	stringInput = strings.TrimSpace(stringInput)
	return
}

// 整数値1つ取得
func intStdin() (num int) {
	stringInput := strStdin()
	num, _ = strconv.Atoi(strings.TrimSpace(stringInput))
	return
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

/* template */

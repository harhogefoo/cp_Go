package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

// 文字列を1行取得
func strStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
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
func splitIntStdin(delim string) (intSlices []int, err error) {
	// 文字列スライスを取得
	stringSplitted := splitStrStdin(" ")

	// 整数スライスに保存
	for i := range stringSplitted {
		var iparam int
		iparam, err = strconv.Atoi(stringSplitted[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}

func main() {
	num, _ := splitIntStdin(" ")

	if num[0] == num[1] {
		fmt.Println(num[2])
	} else if num[0] == num[2] {
		fmt.Println(num[1])
	} else {
		fmt.Println(num[0])
	}
}

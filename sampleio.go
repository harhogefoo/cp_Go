package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
  sc.Scan()
  return sc.Text()
}

func nextInt() int {
  sc.Scan()
  i, e := strconv.Atoi(sc.Text())
  if e != nil {
    panic(e)
  }
  return i
}

func splitWithBlank() {
  // 文字列を空白区切りで分割
  s1 := "1 2 3 4"
  cols := strings.Split(s1, " ")
  fmt.Println(cols)
}

func convertStrToInt() {
  // 文字列をint変換
  s2 := "1234"
  n, _ := strconv.Atoi(s2)
  fmt.Println(n)
}

func convertStrToFloat64() {
// 文字列をfloat64に変換する
  s3 := "12.34"
  f, _ := strconv.ParseFloat(s3, 64)
  fmt.Println(f)
}

func disposeByChar() {
  // 文字列1字ずつにアクセス
  s4 := "abcd"
  cols := strings.Split(s4, "")
  for _, c := range cols {
    fmt.Println(c)
  }
}

func extractStr(from int, to int) {
  s1 := "0123456"
  s2 := s1[from:to]
  fmt.Println(s2)
}

func main() {
  line := nextLine()
  fmt.Println(line)

  splitWithBlank()
  convertStrToInt()
  convertStrToFloat64()
  disposeByChar()
  extractStr(2, 4)
}

type Entry struct {
  name string
  value int
}
type List []Entry

func (l List) Len() int {
  return len(l)
}

func (l List) Swap(i, j int) {
  l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
  // 値が等しい場合はキーで昇順にソート
  if l[i].value == l[j].value {
    return l[i].name < l[j].name
  } else {
    return l[i].value < l[j].value
  }
}



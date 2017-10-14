package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "sort"
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

func main() {
  numDict := map[string]int{}
  numDict["A"] = nextInt()
  numDict["B"] = nextInt()
  numDict["C"] = nextInt()

  a := List{}
  for k, v := range numDict {
    e := Entry{k, v}
    a = append(a, e)
  }

  sort.Sort(a)
  for _, e := range a {
    fmt.Println(e.name)
  }
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
    return (l[i].name < l[j].name)
  } else {
    return (l[i].value > l[j].value)
  }
}


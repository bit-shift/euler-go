package main

import (
  "fmt"
  "os"
  )

var solveFuncs = map[string]func() int {
  "1": solve1,
  "5": solve5,
}

func main() {
  for _, problem := range os.Args[1:] {
    if f, ok := solveFuncs[problem]; ok {
      fmt.Println(problem, "->\t", f())
    } else {
      fmt.Println(problem, "is unsolved or not a valid problem name.")
    }
  }
}

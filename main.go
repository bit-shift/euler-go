package main

import (
  "fmt"
  "os"
  )

var solveFuncs = map[string]func() int {
  "1": func() (sum int) {
    /* all multiples of 3 or 5 are at certain offsets from multiples of 15,
       since 15 is their common multiple */
    offsets := []int{3, 5, 6, 9, 10, 12, 15}

    for i := 0; i < 1000; i += 15 {
      for _, offset := range offsets {
        if i + offset < 1000 {
          sum += (i + offset)
        }
      }
    }

    return
  },
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

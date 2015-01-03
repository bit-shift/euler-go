package main

func solve1() (sum int) {
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
}

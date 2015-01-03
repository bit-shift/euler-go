package main

func generateUints(out chan<- uint) {
  for i := uint(2); ; i++ {
    out <- i
  }
}

func filterDivisible(in <-chan uint, out chan<- uint, divisor uint) {
  for {
    i := <-in
    if i % divisor != 0 {
      out <- i
    }
  }
}

func generatePrimes(out chan<- uint) {
  ch := make(chan uint)
  go generateUints(ch)
  for {
    prime := <-ch
    out <- prime

    ch1 := make(chan uint)
    go filterDivisible(ch, ch1, prime)
    ch = ch1
  }
}

func uintSum(xs []uint) (sum uint) {
  for _, x := range xs {
    sum += x
  }

  return
}

func uintLcm(xs []uint) (lcm uint) {
  lcm = 1
  divisor := uint(2)
  lenXs := uint(len(xs))

  primes := make(chan uint)
  go generatePrimes(primes)

  for uintSum(xs) != lenXs {  // sum == length only when everything is 1, as long as all xs began >=1
    divisionSucceeded := false

    for i, x := range xs {
      if x % divisor == 0 {
        xs[i] = x / divisor
        divisionSucceeded = true
      }
    }

    // if at least one number divided by this divisor, multiply it into the lcm
    if divisionSucceeded {
      lcm *= divisor
    } else {
      divisor = <-primes  // use a prime sieve to speed this up for large values of x
    }
  }

  return
}

package day01

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "aoc/2020/pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 1)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  return findTwoFactors(2020).Product()
}

func Puzzle2() int {
  return findThreeFactors(2020).Product()
}


// ========== PRIVATE FNS =================================

func data () pie.Ints {
  return reader.Lines("./data/day01/input.txt").Ints()
}

func findTwoFactors (sum int) pie.Ints {
  factors := pie.Ints(make([]int, 0))
  list    := data()
  f0      := 0

  for factors.Product() == 0 {
    f0, list = list.Shift()
    f1 := sum - f0
    if list.Contains(f1) {
      factors = append(factors, f0, f1)
    }
  }

  return factors
}

func findThreeFactors (sum int) pie.Ints {
  factors := pie.Ints(make([]int, 0))
  list    := data()
  f0      := 0

  for factors.Product() == 0 {
    f0, list = list.Shift()
    for index, f1 := range list {
      sublist := list[index:]
      f2      := sum - f0 - f1

      if sublist.Contains(f2) {
        factors = append(factors, f0, f1, f2)
        break
      }
    }
  }

  return factors
}

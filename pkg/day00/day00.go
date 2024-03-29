package day00

import (
  "fmt"

  "github.com/elliotchance/pie/pie"

  "aoc/2020/pkg/reader"
)


// ========== PUBLIC FNS ==================================

func Both() {
  fmt.Println(" ")
  fmt.Println("DAY", 0)
  fmt.Println("  Puzzle 1", "=>", Puzzle1())
  fmt.Println("  Puzzle 2", "=>", Puzzle2())
  fmt.Println(" ")
}

func Puzzle1() int {
  return -1
}

func Puzzle2() int {
  return -2
}


// ========== PRIVATE FNS =================================

func data () pie.Strings {
  lines := reader.Lines("./data/day00/input.txt")

  return lines
}

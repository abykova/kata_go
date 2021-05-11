package kata

import (
  "strings"
)

func RepeatStr(repetitions int, value string) string {
  return strings.Repeat(value, repetitions)
}

func RepeatStrLoopFor(repititions int, value string) (result string) {
  for i:=0; i < repititions; i++ {
      result += value
  }
  return
}

func RepeatStrForTwo(repetitions int, value string) string {
  str := ""
  for i:=0; i < repetitions; i++ {
    str+=value
  }
  return str
}

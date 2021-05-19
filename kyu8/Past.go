package kata

import "time"

func Past(h, m, s int) int {
  sh := h * 3600
  sm := m * 60
  sum := sh + sm + s 
  return sum * 1000
}

func PastTwo(h, m, s int) int {
    milliseconds := s
    milliseconds += m * 60
    milliseconds += h * 3600
    milliseconds *= 1000

    return milliseconds
}

func PastThree(h, m, s int) int {
  return int((time.Hour * time.Duration(h) + time.Minute * time.Duration(m) + time.Second * time.Duration(s)) / time.Millisecond)
}

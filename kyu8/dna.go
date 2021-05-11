package kata

import (
  "strings"
)

func DNAtoRNA(dna string) string {
  return strings.Replace(dna, "T", "U", -1)
}

func DNAtoRNAvariableTwo(dna string) string {
  var response string

  for _, letter := range dna{
    if letter == 'T'{
      letter = 'U'
    }
    response+=string(letter)
  }

  return response
}

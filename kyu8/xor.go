package kata

func Xor(a, b bool) bool {
  if ( a && b ) {
      return false
   } else if ( a || b ) {
      return true
   } else if ( !(a && b) ) {
      return false
   } else if ( !(a || b) ) {
      return true
   } else {
      return false
   }
}

func XorTwo(a, b bool) bool {
  return a != b
}

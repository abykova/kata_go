package kata

func CheckForFactor(base int, factor int) bool {
  fac := base % factor
  if fac == 0 {
    return true
  }
  return false
}

func CheckForFactorTwo(base int, factor int) bool {
    return base % factor == 0;
}

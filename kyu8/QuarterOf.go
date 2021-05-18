package kata

func QuarterOf(month int) (result int) {
  
  switch month{
    case 1,2,3:result = 1
    case 4,5,6:result = 2
    case 7,8,9:result = 3
    case 10,11,12:result = 4
  }
  return
}

func QuarterOfTwoVar(month int) int {
  if month <= 3 {
    return 1
  }
  if month == 4 || month == 5 || month == 6 {
   return 2
  }
  if month == 7 || month == 8 || month == 9{
    return 3
  }
  if month >= 10 {
   return 4
  }
  return month
}

package kata

func Grow(arr []int) int{
  m := 1
  for i:=0; i < len(arr); i++ {
    m *= arr[i]
  }
  return m
}

func GrowTwo(arr []int) int{
  x := 1
  for _, v := range arr {
    x *= v 
  }
  return x
}

func myFunc(a []int) {
  if len(a) == 0 {
    fmt.Println("Slice is empty")
    return
  }
  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
  }
}
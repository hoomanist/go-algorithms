package main

import (
  "fmt"
)


func Search(Target int, data []int) int {
  StartIndex := 0
  EndIndex := len(data)-1
  for StartIndex != EndIndex {
    Mid := (StartIndex + EndIndex) / 2
    if data[Mid] == Target {
      return Mid
    }else if data[Mid] > Target {
      EndIndex = Mid
    }else if data[Mid] < Target {
      StartIndex = Mid
    }
  }
  return -1
}

func main(){
  data := []int{1, 2, 3, 4, 5, 6, 7}
  res := Search(4, data)
  fmt.Println(res)
}

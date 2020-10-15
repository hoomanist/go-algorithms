package main

import (
  "fmt"
  "os"
)

func Search(Target int, data []int) (int, error) {
  StartIndex := 0
  EndIndex := len(data)-1
  for StartIndex != EndIndex {
    Mid := (StartIndex + EndIndex) / 2
    if data[Mid] == Target {
      return Mid, nil
    }else if data[Mid] > Target {
      EndIndex = Mid
    }else if data[Mid] < Target {
      StartIndex = Mid
    }
  }
  return -1, fmt.Errorf("No such number")
}

func main(){
  data := []int{1, 2, 3, 4, 5, 6, 7}
  res, err := Search(6, data)
  if err != nil {
    fmt.Println("not found")
    os.Exit(0)
  }
  fmt.Println(res)
}

package main

import "fmt"

func main() {
  nums := []int{1,2,3}
  fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
  set := map[int]bool{}
  for _,num := range nums {
    if(set[num]==true){
      return true
    }
    set[num] = true
  }
  return false
}

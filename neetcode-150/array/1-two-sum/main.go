package main

import "fmt"

func main() {
  nums := []int{2,7,11,15}
  target := 9
  fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
  numMap := map[int]int{}
  result := []int{}
  for i:=0; i< len(nums);i++ {
    diff := target - nums[i]
    value, ok := numMap[diff]
    if(ok) {
      result =append(result, value)
      result = append(result, i)
      return  result
    } else {
      numMap[nums[i]] = i
    }
  }
  return result
}

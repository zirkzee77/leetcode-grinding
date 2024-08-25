package main

import "fmt"

func main() {
  s := "(("
  fmt.Println(isValid(s))
}

func isValid(s string)bool {
  length := len(s)
  if (length%2!=0) {
    return false
  }
  stack := []rune{}
  pMap := map[rune]rune{
    '(': ')',
    '{': '}',
    '[': ']',
  }
  for _, char := range s {
    fmt.Println(stack)
    _, ok := pMap[char]
    if (ok) {
      stack = append(stack, char)
      continue
    } else if (len(stack)>0 && char == pMap[stack[len(stack)-1]]) {
      stack = stack[:len(stack)-1]
      continue
    } else {
      return false
    }
  }
  return len(stack)==0
}

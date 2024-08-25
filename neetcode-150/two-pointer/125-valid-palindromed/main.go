package main

import "fmt"

func main() {
  s := "addsagda"
  fmt.Print(isPalindrome(s))
}

func isPalindrome(s string) bool {
  runeSlice := []rune(s)
  i := 0
  j := len(runeSlice)-1
  for i<j {
    left := convertRune(runeSlice[i])
    right := convertRune(runeSlice[j])
      fmt.Println(runeSlice[i])
      fmt.Println(runeSlice[j])
    fmt.Println(left)
    fmt.Println(right)
    if (left ==0) {
      i++
      continue
    }
    if (right ==0) {
      j--
      continue
    }
    if (left == right) {
      i++
      j--
      continue
    } else {
      return false
    } 
  }
  return true

}

func convertRune(r rune) rune {
  if (97<= r && r <=122) {
    return r
  } else if (48<=r && r<=57){
    return r
  } else if (65<= r && r <=90) {
    return r+32
  }else {
    return 0
  }
}


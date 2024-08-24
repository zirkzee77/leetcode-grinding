package main

import "fmt"

func main() {
  firstString := "tar"
  secondString := "rat"
  result := validAnagram2(firstString, secondString)
  fmt.Println(result)
}

func validAnagram(s string, t string) bool {
  if (len(s) != len(t)) {
    return false
  }
  runeSet := map[rune]int{}
  for _,char := range s {
    if (runeSet[char] > 0) {
      runeSet[char] = runeSet[char] +1
    } else {
      runeSet[char] = 1
    }
  }
  for _,char := range t {
    if (runeSet[char] >0) {
      runeSet[char] = runeSet[char] -1
    } else {
    return false
    }
  }
  return true
}

func validAnagram2(s string, t string) bool {
  if (len(s) != len(t)) {
    return false
  }
  runeSet := map[rune]int{}
  for _,char := range s {
    runeSet[char]++
  }
  for _,char := range t {
    if (runeSet[char] > 0) {
      runeSet[char]--
    } else {
      return false
    }
  }
  return true
}

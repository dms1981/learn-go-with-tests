package main

import (
  "fmt"
)

const repeatCount = 5

func Repeat(character string, iterations int) string {
  var repeated string
  for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}

func main() {
  var character string
  var repetitions int
  fmt.Println("Character to be repeated:")
  fmt.Scan(&character)
  fmt.Println("Number of repetitions to be performed:")
  fmt.Scan(&repetitions)
  fmt.Println(Repeat(character, repetitions))
}
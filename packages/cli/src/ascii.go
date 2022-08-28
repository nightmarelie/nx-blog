package main

import (
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
    fmt.Println("Hello, World!")
    dat, err := os.ReadFile("src/cow.txt")
    check(err)
    fmt.Print(string(dat))
}

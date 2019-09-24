package main

import (
	"fmt"
  "github.com/k0kubun/go-ansi"
  "math/rand"
  "time"
)

// https://www.calhoun.io/creating-random-strings-in-go/
const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func PrintOverwrite(path string) {
		ansi.EraseInLine(2)
		ansi.CursorHorizontalAbsolute(0)
		fmt.Print("- ", path)
}

func main() {
  fmt.Println("Println:")
  fmt.Print("Print:\n")

  fmt.Println("\\r:")
  fmt.Print("line111\rline22\rline3")
  fmt.Println()

  const EL = "\033[K"
  fmt.Println("\\033[K:")
  fmt.Print("line1", EL, "line2", EL, "line3")
  fmt.Println()

  fmt.Println("ansi.EraseInLine:")
  fmt.Print("line111")
  ansi.EraseInLine(2)
  ansi.CursorHorizontalAbsolute(0)
  fmt.Print("line22")
  fmt.Println();

  fmt.Println("loop:")
  for i := 0; i < 10; i++ {
    var randomString = StringWithCharset(seededRand.Intn(25), charset)
    PrintOverwrite(randomString)
    time.Sleep(100 * time.Millisecond)
  }
  fmt.Println()
}

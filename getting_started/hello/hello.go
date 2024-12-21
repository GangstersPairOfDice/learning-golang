package main // declares a main package.

import "fmt" // imports standard lib fmt package, contains funcs for stdio afaik
// standard lib = package which comes a language, without any external downloads


import "rsc.io/quote" // imports external package from pkg.go.dev

// main func impl. to print message to console
func main() {
  fmt.Println(quote.Go())
}

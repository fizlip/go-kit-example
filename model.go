package main

// String interface provided operations on strings
type StringService interface {
  UpperCase(string) (string, error)
  Count(string) int
}



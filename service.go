package main 

import (
  "errors"
  "strings"
)

// String interface provided operations on strings
type StringService interface {
  UpperCase(string) (string, error)
  Count(string) int
}

type stringService struct{}

func (stringService) UpperCase(s string) (string, error){
  if s == ""{
    return "", ErrEmpty
  }
  return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
  return len(s)
}

// Return error when string is empty
var ErrEmpty = errors.New("Empty string")

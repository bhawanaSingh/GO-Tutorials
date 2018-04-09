package main

import "fmt"

func main() {
  for i := 1; i<=10; i++{
  fmt.Println(i)
  }
  
  var j int = 10
  for j>= 1  {
  fmt.Println(j) 
  j--
  }
  
  for k:= 1 ; k<=20 ; k++ {
    if k%2 == 0 {
      fmt.Println("even number", k)
    } else {
      fmt.Println("odd number", k)
    }
  } 
}

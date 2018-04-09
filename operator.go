package main
import "fmt"
func main() {
      //Arithmetic operator
             var a, b int = 20, 11
             fmt.Println(a+b, a-b, a*b , a%b)
  
     //Relational operator
             var x, y  int= 20, 30
             fmt.Println("x==y", x==y,"\n", "x<y" ,x<y, "\n", "x>y", x>y , "\n", "x!=y",x!=y )
  
    // Logical operator
         var T bool = true
         var F bool = false
         fmt.Println( T&&F, T||F, !(T&&F))
  
   // Bitwise operator
         var d, r int = 60, 13
         fmt.Println(d&r, d|r, a<<2, a>>2, a^b)
 
      fmt.Println(&a)// to show address of variable "a"
       var p *int // how to declare pointer variable, both a and p show the same address
       p = &a
       fmt.Println(p)
}
